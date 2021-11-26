package ormindex

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/orm/model/ormiterator"

	"github.com/cosmos/cosmos-sdk/orm/backend/kv"
	"github.com/cosmos/cosmos-sdk/orm/encoding/ormkv"
)

type SingletonIndex struct {
	*ormkv.SingletonKey
}

func (s SingletonIndex) PrefixIterator(store kv.IndexCommitmentReadStore, _ []protoreflect.Value, options IteratorOptions) ormiterator.Iterator {
	return prefixIterator(store.ReadCommitmentStore(), store, s, s.Prefix, options)
}

func (s SingletonIndex) RangeIterator(store kv.IndexCommitmentReadStore, _, _ []protoreflect.Value, options IteratorOptions) ormiterator.Iterator {
	return rangeIterator(store.ReadCommitmentStore(), store, s, s.Prefix, s.Prefix, options)
}

func (s SingletonIndex) doNotImplement() {}

func (s SingletonIndex) Has(store kv.IndexCommitmentReadStore, _ []protoreflect.Value) (found bool, err error) {
	return store.ReadCommitmentStore().Has(s.Prefix)
}

func (s SingletonIndex) Get(store kv.IndexCommitmentReadStore, _ []protoreflect.Value, message proto.Message) (found bool, err error) {
	bz, err := store.ReadCommitmentStore().Get(s.Prefix)
	if err != nil {
		return false, err
	}

	if len(bz) == 0 {
		return false, nil
	}

	return true, proto.Unmarshal(bz, message)
}

func (s SingletonIndex) Fields() []protoreflect.Name {
	return nil
}

func (s SingletonIndex) ReadValueFromIndexKey(_ kv.IndexCommitmentReadStore, _ []protoreflect.Value, value []byte, message proto.Message) error {
	return proto.Unmarshal(value, message)
}

var _ UniqueIndex = &SingletonIndex{}
