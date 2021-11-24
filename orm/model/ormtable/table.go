package ormtable

import (
	"encoding/json"
	"io"

	"github.com/cosmos/cosmos-sdk/orm/encoding/ormkv"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/orm/backend/kv"
	"github.com/cosmos/cosmos-sdk/orm/model/ormindex"
)

type Table interface {
	Save(store kv.Store, message proto.Message, mode SaveMode) error
	Delete(store kv.Store, primaryKey []protoreflect.Value) error

	PrimaryKey() ormindex.UniqueIndex
	GetIndex(fields string) ormindex.Index
	GetUniqueIndex(fields string) ormindex.UniqueIndex
	Indexes() []ormindex.Indexer

	Decode(k []byte, v []byte) (ormkv.Entry, error)

	DefaultJSON() json.RawMessage
	ValidateJSON(io.Reader) error
	ImportJSON(kv.Store, io.Reader) error
	ExportJSON(kv.ReadStore, io.Writer) error
}

type SaveMode int

const (
	SAVE_MODE_DEFAULT SaveMode = iota
	SAVE_MODE_CREATE
	SAVE_MODE_UPDATE
)
