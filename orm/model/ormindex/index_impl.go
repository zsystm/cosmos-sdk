package ormindex

import (
	"bytes"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/orm/backend/kv"
	"github.com/cosmos/cosmos-sdk/orm/encoding/ormkv"
)

type IndexImpl struct {
	*ormkv.IndexKeyCodec
	primaryKey PrimaryKey
}

var _ Indexer = &IndexImpl{}
var _ Index = &IndexImpl{}

var sentinelValue = []byte{0}

func (s IndexImpl) doNotImplement() {}

func (s IndexImpl) Fields() []protoreflect.Name {
	panic("implement me")
}

func (s IndexImpl) OnCreate(store kv.Store, message protoreflect.Message) error {
	_, key, err := s.EncodeFromMessage(message)
	if err != nil {
		return err
	}
	return store.Set(key, sentinelValue)
}

func (s IndexImpl) OnUpdate(store kv.Store, new, existing protoreflect.Message) error {
	newValues := s.GetValues(new)
	existingValues := s.GetValues(existing)
	if s.CompareValues(newValues, existingValues) == 0 {
		return nil
	}

	existingKey, err := s.Encode(existingValues)
	if err != nil {
		return err
	}
	err = store.Delete(existingKey)
	if err != nil {
		return err
	}

	newKey, err := s.Encode(newValues)
	if err != nil {
		return err
	}
	return store.Set(newKey, sentinelValue)
}

func (s IndexImpl) OnDelete(store kv.Store, message protoreflect.Message) error {
	_, key, err := s.EncodeFromMessage(message)
	if err != nil {
		return err
	}
	return store.Delete(key)
}

func (s IndexImpl) PrefixKey(values []protoreflect.Value) ([]byte, error) {
	return s.Encode(values)
}

func (s IndexImpl) ReadValueFromIndexKey(store kv.IndexCommitmentReadStore, key, _ []byte, message proto.Message) error {
	pkValues, err := s.ReadPrimaryKey(bytes.NewReader(key))
	if err != nil {
		return err
	}

	found, err := s.primaryKey.Get(store, pkValues, message)
	if err != nil {
		return err
	}

	if !found {
		return ormerrors.UnexpectedError.Wrapf("can't find primary key")
	}

	return nil
}
