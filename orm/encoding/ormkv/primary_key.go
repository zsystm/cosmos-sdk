package ormkv

import (
	"bytes"

	ormv1alpha1 "github.com/cosmos/cosmos-sdk/api/cosmos/orm/v1alpha1"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"

	"google.golang.org/protobuf/proto"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type PrimaryKeyCodec struct {
	*KeyCodec
	Type protoreflect.MessageType
}

func (p PrimaryKeyCodec) GetIndexValues(k, _ []byte) ([]protoreflect.Value, error) {
	return p.Decode(bytes.NewReader(k))
}

func (p PrimaryKeyCodec) GetPrimaryKeyValues(k, _ []byte) ([]protoreflect.Value, error) {
	return p.Decode(bytes.NewReader(k))
}

var _ IndexCodecI = PrimaryKeyCodec{}

func NewPrimaryKeyCodec(
	prefix []byte,
	messageType protoreflect.MessageType,
	tableDesc *ormv1alpha1.TableDescriptor,
) (*PrimaryKeyCodec, error) {
	tableId := tableDesc.Id
	if tableId == 0 {
		return nil, ormerrors.InvalidTableId.Wrapf("table %s", messageType.Descriptor().FullName())
	}

	primaryKeyDescriptor := tableDesc.PrimaryKey
	if primaryKeyDescriptor == nil {
		return nil, ormerrors.MissingPrimaryKey.Wrap(string(messageType.Descriptor().FullName()))
	}

	desc := messageType.Descriptor()
	pkFields, err := GetFieldDescriptors(desc, primaryKeyDescriptor.Fields)
	if err != nil {
		return nil, err
	}

	if primaryKeyDescriptor.AutoIncrement {
		if len(pkFields) != 1 && pkFields[0].Kind() != protoreflect.Uint64Kind {
			return nil, ormerrors.InvalidAutoIncrementKey.Wrapf("got %s for %s", primaryKeyDescriptor.Fields, desc.FullName())
		}
	}

	pkPrefix := AppendVarUint32(prefix, tableDesc.Id)
	pkPrefix = AppendVarUint32(pkPrefix, 0)

	cdc, err := NewKeyCodec(pkPrefix, pkFields)

	return &PrimaryKeyCodec{
		KeyCodec: cdc,
		Type:     messageType,
	}, nil
}

func (p PrimaryKeyCodec) DecodeKV(k, v []byte) (Entry, error) {
	vals, err := p.Decode(bytes.NewReader(k))
	if err != nil {
		return nil, err
	}

	msg := p.Type.New().Interface()
	err = proto.Unmarshal(v, msg)
	if err != nil {
		return nil, err
	}

	return PrimaryKeyEntry{
		Key:   vals,
		Value: msg,
	}, nil
}

func (p PrimaryKeyCodec) EncodeKV(entry Entry) (k, v []byte, err error) {
	pkEntry, ok := entry.(PrimaryKeyEntry)
	if !ok {
		return nil, nil, ormerrors.BadDecodeEntry
	}

	if pkEntry.Value.ProtoReflect().Descriptor().FullName() != p.Type.Descriptor().FullName() {
		return nil, nil, ormerrors.BadDecodeEntry
	}

	bz, err := p.KeyCodec.Encode(pkEntry.Key)
	if err != nil {
		return nil, nil, err
	}

	v, err = proto.MarshalOptions{Deterministic: true}.Marshal(pkEntry.Value)
	if err != nil {
		return nil, nil, err
	}

	return bz, v, nil
}

func (p *PrimaryKeyCodec) ClearValues(mref protoreflect.Message) {
	for _, f := range p.FieldDescriptors {
		mref.Clear(f)
	}
}

func (p *PrimaryKeyCodec) Unmarshal(key []protoreflect.Value, value []byte, message proto.Message) error {
	err := proto.Unmarshal(value, message)
	if err != nil {
		return err
	}

	// rehydrate primary key
	p.SetValues(message.ProtoReflect(), key)
	return nil
}
