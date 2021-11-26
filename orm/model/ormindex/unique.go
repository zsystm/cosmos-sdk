package ormindex

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"

	"github.com/cosmos/cosmos-sdk/orm/model/ormiterator"

	"github.com/cosmos/cosmos-sdk/orm/backend/kv"
	"github.com/cosmos/cosmos-sdk/orm/encoding/ormkv"
)

type UniqueIndexImpl struct {
	ormkv.UniqueKeyCodec
	primaryKey PrimaryKey
}

func (u UniqueIndexImpl) PrefixIterator(store kv.IndexCommitmentReadStore, prefix []protoreflect.Value, options IteratorOptions) ormiterator.Iterator {
	prefixBz, err := u.KeyCodec.Encode(prefix)
	if err != nil {
		return ormiterator.ErrIterator{Err: err}
	}

	return prefixIterator(store.ReadIndexStore(), store, u, prefixBz, options)
}

func (u UniqueIndexImpl) RangeIterator(store kv.IndexCommitmentReadStore, start, end []protoreflect.Value, options IteratorOptions) ormiterator.Iterator {
	err := u.KeyCodec.CheckValidRangeIterationKeys(start, end)
	if err != nil {
		return ormiterator.ErrIterator{Err: err}
	}

	startBz, err := u.KeyCodec.Encode(start)
	if err != nil {
		return ormiterator.ErrIterator{Err: err}
	}

	endBz, err := u.KeyCodec.Encode(end)
	if err != nil {
		return ormiterator.ErrIterator{Err: err}
	}

	return rangeIterator(store.ReadIndexStore(), store, u, startBz, endBz, options)
}

func (u UniqueIndexImpl) doNotImplement() {}

func (u UniqueIndexImpl) Fields() []protoreflect.Name {
	return u.KeyCodec.FieldNames
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

func (u UniqueIndexImpl) ReadValueFromIndexKey(store kv.IndexCommitmentReadStore, key, value []byte, message proto.Message) error {
	pk, err := u.ExtractPrimaryKey(key, value)
	if err != nil {
		return err
	}

	found, err := u.primaryKey.Get(store, pk, message)
	if err != nil {
		return err
	}

	if !found {
		return ormerrors.UnexpectedError.Wrapf("can't find primary key")
	}

	return nil
}

var _ Indexer = &UniqueIndexImpl{}
var _ UniqueIndex = &UniqueIndexImpl{}
