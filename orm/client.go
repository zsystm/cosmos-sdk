package orm

import (
	"google.golang.org/protobuf/proto"

	"github.com/cosmos/cosmos-sdk/orm/encoding/encodeutil"
	"github.com/cosmos/cosmos-sdk/orm/model/kvstore"
	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
)

type ReadClient struct {
	Schema      Schema
	ReadBackend kvstore.ReadBackend
}

type Client struct {
	*ReadClient
	Backend kvstore.Backend
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

	return index.Get(c.ReadBackend, encodeutil.ValuesOf(fields...), message)
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

	return index.Has(c.ReadBackend, encodeutil.ValuesOf(fields...))
}

func (c Client) Save(message proto.Message) error {
	table, err := c.Schema.GetTable(message)
	if err != nil {
		return err
	}

	return table.Save(c.Backend, message, ormtable.SAVE_MODE_DEFAULT)
}

func (c Client) Insert(message proto.Message) error {
	table, err := c.Schema.GetTable(message)
	if err != nil {
		return err
	}

	return table.Save(c.Backend, message, ormtable.SAVE_MODE_INSERT)
}

func (c Client) Update(message proto.Message) error {
	table, err := c.Schema.GetTable(message)
	if err != nil {
		return err
	}

	return table.Save(c.Backend, message, ormtable.SAVE_MODE_UPDATE)
}

func (c Client) Delete(message proto.Message) error {
	table, err := c.Schema.GetTable(message)
	if err != nil {
		return err
	}

	return table.DeleteMessage(c.Backend, message)
}

func (c ReadClient) List(message proto.Message, options ...ormlist.Option) (ormtable.Iterator, error) {
	table, err := c.Schema.GetTable(message)
	if err != nil {
		return nil, err
	}

	return ormlist.Iterator(c.ReadBackend, table, options...)
}
