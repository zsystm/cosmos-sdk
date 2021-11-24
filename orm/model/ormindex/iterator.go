package ormindex

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/orm/backend/kv"
	"github.com/cosmos/cosmos-sdk/orm/model/ormiterator"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

func PrefixIterator(store kv.ReadStore, index Index, prefix []protoreflect.Value, options IteratorOptions) ormiterator.Iterator {
	prefixBz, err := index.PrefixKey(prefix)
	if err != nil {
		return ormiterator.ErrIterator{Err: err}
	}

	return iterator(store, index, prefixBz, prefixBz, options)
}

func RangeIterator(store kv.ReadStore, index Index, start, end []protoreflect.Value, options IteratorOptions) ormiterator.Iterator {
	startBz, err := index.PrefixKey(start)
	if err != nil {
		return ormiterator.ErrIterator{Err: err}
	}

	endBz, err := index.PrefixKey(end)
	if err != nil {
		return ormiterator.ErrIterator{Err: err}
	}

	return iterator(store, index, startBz, endBz, options)
}

func iterator(store kv.ReadStore, index Index, start, end []byte, options IteratorOptions) ormiterator.Iterator {
	if !options.Reverse {
		if len(options.Cursor) != 0 {
			start = options.Cursor
		}
		it := store.Iterator(start, storetypes.PrefixEndBytes(end))
		return &indexIterator{
			index:    nil,
			store:    nil,
			iterator: it,
			started:  false,
		}
	} else {
		if len(options.Cursor) != 0 {
			end = options.Cursor
		}
		it := store.ReverseIterator(start, storetypes.PrefixEndBytes(end))
		return &indexIterator{
			index:    nil,
			store:    nil,
			iterator: it,
			started:  false,
		}
	}
}

type indexIterator struct {
	ormiterator.UnimplementedIterator

	index    Index
	store    kv.ReadStore
	iterator kv.Iterator
	started  bool
}

func (i indexIterator) Next(message proto.Message) (bool, error) {
	if !i.started {
		i.started = true
	} else {
		i.iterator.Next()
	}

	if !i.iterator.Valid() {
		return false, nil
	}

	err := i.index.ReadValueFromIndexKey(i.store, i.iterator.Key(), i.iterator.Value(), message)
	return true, err
}

func (i indexIterator) Cursor() ormiterator.Cursor {
	return i.iterator.Key()
}

func (i indexIterator) Close() {
	_ = i.iterator.Close()
}

var _ ormiterator.Iterator = &indexIterator{}
