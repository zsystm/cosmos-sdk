package list

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"

	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/orm/encoding/ormkv"
	"github.com/cosmos/cosmos-sdk/orm/model/kvstore"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
)

type Option interface {
	applyListOption(*opts)
}

type option func(*opts)

func (o option) applyListOption(opts *opts) { o(opts) }

func UseIndex(fieldNames ormtable.FieldNames) Option {
	return option(func(o *opts) {
		o.index = fieldNames
	})
}

func Prefix(values ...interface{}) Option {
	return option(func(o *opts) {
		o.prefix = ormkv.ValuesOf(values)
	})
}

func Start(values ...interface{}) Option {
	return option(func(o *opts) {
		o.start = ormkv.ValuesOf(values)
	})
}

func End(values ...interface{}) Option {
	return option(func(o *opts) {
		o.end = ormkv.ValuesOf(values)
	})
}

func Reverse() Option {
	return option(func(o *opts) {
		o.reverse = true
	})
}

func WithCursor(cursor ormtable.Cursor) Option {
	return option(func(o *opts) {
		o.cursor = cursor
	})
}

type opts struct {
	index   ormtable.FieldNames
	reverse bool
	cursor  ormtable.Cursor
	prefix  []protoreflect.Value
	start   []protoreflect.Value
	end     []protoreflect.Value
}

func buildOpts(options ...Option) *opts {
	o := &opts{}
	for _, opt := range options {
		opt.applyListOption(o)
	}
	return o
}

func getIndex(table ormtable.Table, o *opts) (ormtable.Index, error) {
	var index ormtable.Index = table
	if o.index.String() != "" {
		index = table.GetIndex(o.index)
		if index == nil {
			return nil, ormerrors.CantFindIndex.Wrapf(
				"for table %s with fields %s",
				table.MessageType().Descriptor().FullName(),
				o.index,
			)
		}
	}
	return index, nil
}

func Iterator(store kvstore.IndexCommitmentReadStore, table ormtable.Table, options ...Option) (ormtable.Iterator, error) {
	o := buildOpts(options...)

	index, err := getIndex(table, o)
	if err != nil {
		return nil, err
	}

	itOpts := ormtable.IteratorOptions{
		Reverse: o.reverse,
		Cursor:  o.cursor,
	}

	if o.start != nil || o.end != nil {
		if o.prefix != nil {
			return nil, fmt.Errorf("can either use Start/End or Prefix, not both")
		}
		return index.RangeIterator(store, o.start, o.end, itOpts)
	} else {
		return index.PrefixIterator(store, o.prefix, itOpts)
	}
}
