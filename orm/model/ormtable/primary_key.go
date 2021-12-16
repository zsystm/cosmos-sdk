package ormtable

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/encoding/encodeutil"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/orm/encoding/ormkv"
)

// PrimaryKeyIndex defines an UniqueIndex for the primary key.
type PrimaryKeyIndex struct {
	*ormkv.PrimaryKeyCodec
	getReadContext func(context.Context) (ReadContext, error)
}

// NewPrimaryKeyIndex returns a new PrimaryKeyIndex.
func NewPrimaryKeyIndex(primaryKeyCodec *ormkv.PrimaryKeyCodec) *PrimaryKeyIndex {
	return &PrimaryKeyIndex{PrimaryKeyCodec: primaryKeyCodec}
}

func (p PrimaryKeyIndex) PrefixIterator(store ReadContext, prefix []protoreflect.Value, options IteratorOptions) (Iterator, error) {
	prefixBz, err := p.EncodeKey(prefix)
	if err != nil {
		return nil, err
	}

	return prefixIterator(store.CommitmentStoreReader(), store, p, prefixBz, options)
}

func (p PrimaryKeyIndex) RangeIterator(store ReadContext, start, end []protoreflect.Value, options IteratorOptions) (Iterator, error) {
	err := p.CheckValidRangeIterationKeys(start, end)
	if err != nil {
		return nil, err
	}

	startBz, err := p.EncodeKey(start)
	if err != nil {
		return nil, err
	}

	endBz, err := p.EncodeKey(end)
	if err != nil {
		return nil, err
	}

	fullEndKey := len(p.GetFieldNames()) == len(end)

	return rangeIterator(store.CommitmentStoreReader(), store, p, startBz, endBz, fullEndKey, options)
}

func (p PrimaryKeyIndex) doNotImplement() {}

func (p PrimaryKeyIndex) Has(context context.Context, key ...interface{}) (found bool, err error) {
	ctx, err := p.getReadContext(context)
	if err != nil {
		return false, err
	}

	keyBz, err := p.EncodeKey(encodeutil.ValuesOf(key...))
	if err != nil {
		return false, err
	}

	return ctx.CommitmentStoreReader().Has(keyBz)
}

func (p PrimaryKeyIndex) Get(store ReadContext, keyValues []protoreflect.Value, message proto.Message) (found bool, err error) {
	key, err := p.EncodeKey(keyValues)
	if err != nil {
		return false, err
	}

	return p.GetByKeyBytes(store, key, keyValues, message)
}

func (p PrimaryKeyIndex) GetByKeyBytes(store ReadContext, key []byte, keyValues []protoreflect.Value, message proto.Message) (found bool, err error) {
	bz, err := store.CommitmentStoreReader().Get(key)
	if err != nil {
		return false, err
	}

	if bz == nil {
		return false, nil
	}

	return true, p.Unmarshal(keyValues, bz, message)
}

func (p PrimaryKeyIndex) readValueFromIndexKey(_ ReadContext, primaryKey []protoreflect.Value, value []byte, message proto.Message) error {
	return p.Unmarshal(primaryKey, value, message)
}

var _ UniqueIndex = &PrimaryKeyIndex{}
