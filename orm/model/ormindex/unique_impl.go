package ormindex

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/orm/backend/kv"
	"github.com/cosmos/cosmos-sdk/orm/encoding/ormkv"
)

type UniqueIndexImpl struct {
	ormkv.UniqueKeyCodec
}

var _ Indexer = &UniqueIndexImpl{}
var _ UniqueIndex = &UniqueIndexImpl{}

func (u UniqueIndexImpl) doNotImplement() {}

func (u UniqueIndexImpl) Fields() []protoreflect.Name {
	panic("implement me")
}

func (u UniqueIndexImpl) Has(store kv.IndexCommitmentReadStore, keyValues []protoreflect.Value) (found bool, err error) {
	key, err := u.KeyCodec.Encode(keyValues)
	if err != nil {
		return false, err
	}

	return store.ReadIndexStore().Has(key)
}

func (u UniqueIndexImpl) Get(store kv.IndexCommitmentReadStore, keyValues []protoreflect.Value, message proto.Message) (found bool, err error) {
	key, err := u.KeyCodec.Encode(keyValues)
	if err != nil {
		return false, err
	}

	bz, err := store.ReadIndexStore().Get(key)
	if err != nil {
		return false, err
	}

	if len(bz) == 0 {
		return false, nil
	}

	return true, proto.Unmarshal(bz, message)
}

func (u UniqueIndexImpl) OnCreate(store kv.Store, message protoreflect.Message) error {
	_, key, err := u.KeyCodec.EncodeFromMessage(message)
	if err != nil {
		return err
	}

	_, value, err := u.ValueCodec.EncodeFromMessage(message)
	if err != nil {
		return err
	}

	return store.Set(key, value)
}

func (u UniqueIndexImpl) OnUpdate(store kv.Store, new, existing protoreflect.Message) error {
	newValues := u.KeyCodec.GetValues(new)
	existingValues := u.KeyCodec.GetValues(existing)
	if u.KeyCodec.CompareValues(newValues, existingValues) == 0 {
		return nil
	}

	existingKey, err := u.KeyCodec.Encode(existingValues)
	if err != nil {
		return err
	}
	err = store.Delete(existingKey)
	if err != nil {
		return err
	}

	newKey, err := u.KeyCodec.Encode(newValues)
	if err != nil {
		return err
	}

	_, value, err := u.ValueCodec.EncodeFromMessage(new)
	if err != nil {
		return err
	}

	return store.Set(newKey, value)
}

func (u UniqueIndexImpl) OnDelete(store kv.Store, message protoreflect.Message) error {
	_, key, err := u.KeyCodec.EncodeFromMessage(message)
	if err != nil {
		return err
	}

	_, value, err := u.ValueCodec.EncodeFromMessage(message)
	if err != nil {
		return err
	}

	return store.Set(key, value)
}

func (u UniqueIndexImpl) PrefixKey(values []protoreflect.Value) ([]byte, error) {
	return u.KeyCodec.Encode(values)
}

func (u UniqueIndexImpl) ReadValueFromIndexKey(store kv.IndexCommitmentReadStore, key, value []byte, message proto.Message) error {
	panic("implement me")
}
