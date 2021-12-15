package orm

import (
	"google.golang.org/protobuf/proto"

	"github.com/cosmos/cosmos-sdk/orm/client/list"
	"github.com/cosmos/cosmos-sdk/orm/encoding/ormkv"
	"github.com/cosmos/cosmos-sdk/orm/model/kvstore"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
)

type ReadClient struct {
	Schema Schema
	Store  kvstore.IndexCommitmentReadStore
}

type Client struct {
	*ReadClient
	Store kvstore.IndexCommitmentStore
}

type Schema interface {
	GetTable(message proto.Message) (ormtable.Table, error)
}

func (c ReadClient) Get(message proto.Message, fieldNames ormtable.FieldNames, fields ...interface{}) (found bool, err error) {
	table, err := c.Schema.GetTable(message)
	if err != nil {
		return false, err
	}

	index := table.GetUniqueIndex(fieldNames)
	if index == nil {
		return false, ormerrors.CantFindIndex.Wrapf(
			"can't find unique index on table %s for fields %s",
			message.ProtoReflect().Descriptor().FullName(),
			fieldNames,
		)
	}

	return index.Get(c.Store, ormkv.ValuesOf(fields...), message)
}

func (c ReadClient) Has(message proto.Message, fieldNames ormtable.FieldNames, fields ...interface{}) (found bool, err error) {
	table, err := c.Schema.GetTable(message)
	if err != nil {
		return false, err
	}

	index := table.GetUniqueIndex(fieldNames)
	if index == nil {
		return false, ormerrors.CantFindIndex.Wrapf(
			"can't find unique index on table %s for fields %s",
			message.ProtoReflect().Descriptor().FullName(),
			fieldNames,
		)
	}

	return index.Has(c.Store, ormkv.ValuesOf(fields...))
}

func (c Client) Save(message proto.Message) error {
	table, err := c.Schema.GetTable(message)
	if err != nil {
		return err
	}

	return table.Save(c.Store, message, ormtable.SAVE_MODE_DEFAULT)
}

func (c Client) Insert(message proto.Message) error {
	table, err := c.Schema.GetTable(message)
	if err != nil {
		return err
	}

	return table.Save(c.Store, message, ormtable.SAVE_MODE_INSERT)
}

func (c Client) Update(message proto.Message) error {
	table, err := c.Schema.GetTable(message)
	if err != nil {
		return err
	}

	return table.Save(c.Store, message, ormtable.SAVE_MODE_UPDATE)
}

func (c Client) Delete(message proto.Message) error {
	table, err := c.Schema.GetTable(message)
	if err != nil {
		return err
	}

	return table.DeleteMessage(c.Store, message)
}

func (c ReadClient) List(message proto.Message, options ...list.Option) (ormtable.Iterator, error) {
	table, err := c.Schema.GetTable(message)
	if err != nil {
		return nil, err
	}

	return list.Iterator(c.Store, table, options...)
}
