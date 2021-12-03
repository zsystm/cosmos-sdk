package ormsql

import (
	"fmt"
	"reflect"
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"

	ormv1alpha1 "github.com/cosmos/cosmos-sdk/api/cosmos/orm/v1alpha1"
)

type messageCodec struct {
	msgType     protoreflect.MessageType
	structType  reflect.Type
	fieldCodecs []*fieldCodec
}

func (b *builder) makeMessageCodec(messageType protoreflect.MessageType, tableDesc *ormv1alpha1.TableDescriptor) (*messageCodec, error) {
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

	return &messageCodec{
		msgType:     messageType,
		fieldCodecs: fieldCodecs,
		structType:  reflect.StructOf(structFields),
	}, nil
}

func (m messageCodec) encode(message protoreflect.Message) reflect.Value {
	ptr := reflect.New(m.structType)
	val := ptr.Elem()
	for _, codec := range m.fieldCodecs {
		codec.encode(message, val)
	}
	return ptr
}
