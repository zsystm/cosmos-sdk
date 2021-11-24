package ormtable

import (
	"encoding/json"
	io "io"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/orm/backend/kv"
	"github.com/cosmos/cosmos-sdk/orm/encoding/ormkv"
	"github.com/cosmos/cosmos-sdk/orm/model/ormindex"
)

type TableImpl struct {
	msgType         protoreflect.MessageType
	primaryKey      ormindex.PrimaryKey
	indexes         []ormindex.Index
	indexers        []ormindex.Indexer
	indexesByFields map[string]ormindex.Index
	indexesById     map[uint32]ormindex.Index
	prefix          []byte
}

func (t TableImpl) Save(store kv.Store, message proto.Message, mode SaveMode) error {
	panic("implement me")
}

func (t TableImpl) Delete(store kv.Store, primaryKey []protoreflect.Value) error {
	panic("implement me")
}

func (t TableImpl) PrimaryKey() ormindex.UniqueIndex {
	panic("implement me")
}

func (t TableImpl) GetIndex(fields string) ormindex.Index {
	panic("implement me")
}

func (t TableImpl) GetUniqueIndex(fields string) ormindex.UniqueIndex {
	panic("implement me")
}

func (t TableImpl) Indexes() []ormindex.Indexer {
	panic("implement me")
}

func (t TableImpl) Decode(k []byte, v []byte) (ormkv.Entry, error) {
	panic("implement me")
}

func (t TableImpl) DefaultJSON() json.RawMessage {
	panic("implement me")
}

func (t TableImpl) ValidateJSON(reader io.Reader) error {
	panic("implement me")
}

func (t TableImpl) ImportJSON(store kv.Store, reader io.Reader) error {
	panic("implement me")
}

func (t TableImpl) ExportJSON(store kv.ReadStore, writer io.Writer) error {
	panic("implement me")
}

var _ Table = &TableImpl{}
