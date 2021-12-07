package ormsql

import (
	"fmt"
	"reflect"
	"strings"

	"google.golang.org/protobuf/proto"

	"google.golang.org/protobuf/reflect/protoreflect"

	ormv1alpha1 "github.com/cosmos/cosmos-sdk/api/cosmos/orm/v1alpha1"
)

type messageCodec struct {
	tableName   string
	msgType     protoreflect.MessageType
	structType  reflect.Type
	fieldCodecs []*fieldCodec
}

func (b *schema) getMessageCodec(message proto.Message) (*messageCodec, error) {
	return b.messageCodecForType(message.ProtoReflect().Type())
}

func (b *schema) messageCodecForType(messageType protoreflect.MessageType) (*messageCodec, error) {
	if existing, ok := b.messageCodecs[messageType.Descriptor().FullName()]; ok {
		return existing, nil
	}

	tableDesc := proto.GetExtension(messageType.Descriptor().Options(), ormv1alpha1.E_Table).(*ormv1alpha1.TableDescriptor)
	return b.makeMessageCodec(messageType, tableDesc)
}

func (b *schema) makeMessageCodec(messageType protoreflect.MessageType, tableDesc *ormv1alpha1.TableDescriptor) (*messageCodec, error) {
	if tableDesc.PrimaryKey == nil {
		return nil, fmt.Errorf("missing primary key")
	}

	pk := tableDesc.PrimaryKey
	pkFields := strings.Split(pk.Fields, ",")
	if len(pkFields) == 0 {
		return nil, fmt.Errorf("missing primary key fields")
	}
	pkFieldMap := map[string]bool{}
	for _, k := range pkFields {
		pkFieldMap[k] = true
	}

	desc := messageType.Descriptor()
	fieldDescriptors := desc.Fields()
	n := fieldDescriptors.Len()
	var fieldCodecs []*fieldCodec
	var structFields []reflect.StructField
	for i := 0; i < n; i++ {
		field := fieldDescriptors.Get(i)
		fieldCodec, err := b.makeFieldCodec(field, pkFieldMap[string(field.Name())])
		if err != nil {
			// TODO: return nil, err
			// for now:
			continue
		}
		fieldCodecs = append(fieldCodecs, fieldCodec)
		structFields = append(structFields, fieldCodec.structField)
	}

	tableName := strings.ReplaceAll(string(messageType.Descriptor().FullName()), ".", "_")

	msgCdc := &messageCodec{
		tableName:   tableName,
		msgType:     messageType,
		fieldCodecs: fieldCodecs,
		structType:  reflect.StructOf(structFields),
	}

	b.messageCodecs[messageType.Descriptor().FullName()] = msgCdc
	return msgCdc, nil
}

func (m *messageCodec) encode(message protoreflect.Message) (reflect.Value, error) {
	ptr := reflect.New(m.structType)
	val := ptr.Elem()
	for _, codec := range m.fieldCodecs {
		err := codec.encode(message, val)
		if err != nil {
			return reflect.Value{}, err
		}
	}
	return ptr, nil
}

func (m messageCodec) decode(value reflect.Value, msg protoreflect.Message) error {
	for _, codec := range m.fieldCodecs {
		err := codec.decode(value, msg)
		if err != nil {
			return err
		}
	}
	return nil
}
