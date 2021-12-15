package ormlist

import (
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
)

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
