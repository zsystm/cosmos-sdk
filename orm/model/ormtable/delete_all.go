package ormtable

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"

	"google.golang.org/protobuf/reflect/protoreflect"
)

// DeleteBy deletes all the entries matched by the prefix key.
func DeleteBy(ctx context.Context, table Table, prefixKey ...interface{}) error {
	it, err := table.Iterator(ctx, ormlist.Prefix(prefixKey...))
	if err != nil {
		return err
	}

	return DeleteAll(ctx, table, it)
}

// DeleteAll deletes all the entries in the table matched in the iterator.
func DeleteAll(ctx context.Context, table Table, iterator Iterator) error {
	var pksToDelete [][]protoreflect.Value
	for iterator.Next() {
		_, pk, err := iterator.Keys()
		if err != nil {
			return err
		}
		pksToDelete = append(pksToDelete, pk)
	}
	iterator.Close()

	for _, pk := range pksToDelete {
		vs := make([]interface{}, len(pk))
		for i, value := range pk {
			vs[i] = value.Interface()
		}

		err := table.PrimaryKey().DeleteByKey(ctx, vs...)
		if err != nil {
			return err
		}
	}

	return nil
}
