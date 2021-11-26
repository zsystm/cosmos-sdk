package ormindex

import (
	"google.golang.org/protobuf/proto"

	"github.com/cosmos/cosmos-sdk/orm/backend/kv"
	"github.com/cosmos/cosmos-sdk/orm/model/ormiterator"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

func prefixIterator(iteratorStore kv.ReadStore, store kv.IndexCommitmentReadStore, index Index, prefix []byte, options IteratorOptions) ormiterator.Iterator {
	if !options.Reverse {
		var start []byte
		if len(options.Cursor) != 0 {
			// must start right after cursor
			start = append(options.Cursor, 0x0)
		} else {
			start = prefix
		}
		end := storetypes.PrefixEndBytes(prefix)
		it := iteratorStore.Iterator(start, end)
		return &indexIterator{
			index:    index,
			store:    store,
			iterator: it,
			started:  false,
		}
	} else {
		var end []byte
		if len(options.Cursor) != 0 {
			// end bytes is already exclusive by default
			end = options.Cursor
		} else {
			end = storetypes.PrefixEndBytes(prefix)
		}
		it := iteratorStore.ReverseIterator(prefix, end)
		return &indexIterator{
			index:    index,
			store:    store,
			iterator: it,
			started:  false,
		}
	}
}

func rangeIterator(iteratorStore kv.ReadStore, store kv.IndexCommitmentReadStore, index Index, start, end []byte, options IteratorOptions) ormiterator.Iterator {
	if !options.Reverse {
		if len(options.Cursor) != 0 {
			start = append(options.Cursor, 0)
		}
		it := iteratorStore.Iterator(start, storetypes.InclusiveEndBytes(end))
		return &indexIterator{
			index:    index,
			store:    store,
			iterator: it,
			started:  false,
		}
	} else {
		if len(options.Cursor) != 0 {
			end = options.Cursor
		} else {
			end = storetypes.PrefixEndBytes(end)
		}
		it := iteratorStore.ReverseIterator(start, storetypes.InclusiveEndBytes(end))
		return &indexIterator{
			index:    index,
			store:    store,
			iterator: it,
			started:  false,
		}
	}
}

type indexIterator struct {
	ormiterator.UnimplementedIterator

	index    Index
	store    kv.IndexCommitmentReadStore
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
