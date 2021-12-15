package ormlist

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/orm/model/kvstore"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
)

func Iterator(store kvstore.ReadBackend, table ormtable.Table, options ...Option) (ormtable.Iterator, error) {
	o := buildOpts(options...)

	index, err := getIndex(table, o)
	if err != nil {
		return nil, err
	}

	itOpts := ormtable.IteratorOptions{
		Reverse: o.reverse,
		Cursor:  o.cursor,
	}

	var iterator ormtable.Iterator
	if o.start != nil || o.end != nil {
		if o.prefix != nil {
			return nil, fmt.Errorf("can either use Start/End or Prefix, not both")
		}
		iterator, err = index.RangeIterator(store, o.start, o.end, itOpts)
	} else {
		iterator, err = index.PrefixIterator(store, o.prefix, itOpts)
	}
	if err != nil {
		return nil, err
	}

	//if o.filter != nil {
	//	panic("TODO")
	//}

	return iterator, nil
}
