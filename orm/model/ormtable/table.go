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
	ormindex.UniqueIndex

	Save(store kv.IndexCommitmentStore, message proto.Message, mode SaveMode) error
	Delete(store kv.IndexCommitmentStore, primaryKey []protoreflect.Value) error

	GetIndex(fields Fields) ormindex.Index
	GetUniqueIndex(fields Fields) ormindex.UniqueIndex
	Indexes() []ormindex.Index

	Decode(k []byte, v []byte) (ormkv.Entry, error)

	DefaultJSON() json.RawMessage
	ValidateJSON(io.Reader) error
	ImportJSON(kv.IndexCommitmentStore, io.Reader) error
	ExportJSON(kv.IndexCommitmentReadStore, io.Writer) error
}

type SaveMode int

const (
	SAVE_MODE_DEFAULT SaveMode = iota
	SAVE_MODE_CREATE
	SAVE_MODE_UPDATE
)
