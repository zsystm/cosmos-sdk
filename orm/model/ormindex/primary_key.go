package ormindex

import (
	"bytes"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/orm/backend/kv"
	"github.com/cosmos/cosmos-sdk/orm/encoding/ormkv"
)

type PrimaryKey struct {
	*ormkv.PrimaryKeyCodec
}

func (p PrimaryKey) Fields() []protoreflect.Name {
	panic("implement me")
}

func (p PrimaryKey) Has(store kv.ReadStore, key []protoreflect.Value) (found bool, err error) {
	panic("implement me")
}

func (p PrimaryKey) Get(store kv.ReadStore, key []protoreflect.Value, message proto.Message) (found bool, err error) {
	panic("implement me")
}

func (p PrimaryKey) ReadPrimaryKey(store kv.ReadStore, keyValues []protoreflect.Value, message proto.Message) error {
	key, err := p.Encode(keyValues)
	if err != nil {
		return err
	}

	bz, err := store.Get(key)
	if err != nil {
		return err
	}

	return p.unmarshalMessage(keyValues, bz, message)
}

func (p PrimaryKey) unmarshalMessage(keyValues []protoreflect.Value, value []byte, message proto.Message) error {
	err := proto.Unmarshal(value, message)
	if err != nil {
		return err
	}

	// rehydrate primary key
	p.SetValues(message.ProtoReflect(), keyValues)
	return nil
}

func (p PrimaryKey) ReadValueFromIndexKey(store kv.ReadStore, key, value []byte, message proto.Message) error {
	keyValues, err := p.Decode(bytes.NewReader(key))
	if err != nil {
		return err
	}

	return p.unmarshalMessage(keyValues, value, message)
}

func (p PrimaryKey) PrefixKey(values []protoreflect.Value) ([]byte, error) {
	return p.Encode(values)
}

var _ UniqueIndex = &PrimaryKey{}
