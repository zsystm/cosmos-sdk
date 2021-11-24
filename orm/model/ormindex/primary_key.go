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

func (p PrimaryKey) doNotImplement() {}

func (p PrimaryKey) Fields() []protoreflect.Name {
	panic("implement me")
}

func (p PrimaryKey) Has(store kv.IndexCommitmentReadStore, key []protoreflect.Value) (found bool, err error) {
	keyBz, err := p.Encode(key)
	if err != nil {
		return false, err
	}

	return store.ReadCommitmentStore().Has(keyBz)
}

func (p PrimaryKey) Get(store kv.IndexCommitmentReadStore, keyValues []protoreflect.Value, message proto.Message) (found bool, err error) {
	key, err := p.Encode(keyValues)
	if err != nil {
		return false, err
	}

	return p.GetByKeyBytes(store, key, keyValues, message)
}

func (p PrimaryKey) GetByKeyBytes(store kv.IndexCommitmentReadStore, key []byte, keyValues []protoreflect.Value, message proto.Message) (found bool, err error) {
	bz, err := store.ReadCommitmentStore().Get(key)
	if err != nil {
		return true, err
	}

	return true, p.unmarshalMessage(keyValues, bz, message)
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

func (p PrimaryKey) ReadValueFromIndexKey(store kv.IndexCommitmentReadStore, key, value []byte, message proto.Message) error {
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
