package ormschema

import (
	"bytes"
	"encoding/binary"
	"math"

	"github.com/cosmos/cosmos-sdk/orm/encoding/ormkv"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type ModuleSchema struct {
	modulePrefix   []byte
	tablesById     map[uint32]ormtable.Table
	tablesByName   map[protoreflect.FullName]ormtable.Table
	schemaKeyCodec ormkv.SchemaKeyCodec
}

func (m ModuleSchema) DecodeKV(k, v []byte) (ormkv.Entry, error) {
	r := bytes.NewReader(k)
	if bytes.HasPrefix(k, m.modulePrefix) {
		err := ormkv.SkipPrefix(r, m.modulePrefix)
		if err != nil {
			return nil, err
		}

		id, err := binary.ReadUvarint(r)
		if err != nil {
			return nil, err
		}

		// schema sub-store
		if id == 0 {
			schemaKey, err := binary.ReadUvarint(r)
			if err != nil {
				return nil, err
			}

			switch schemaKey {
			case 0:
				return m.schemaKeyCodec.DecodeKV(k, v)
			default:
				return nil, ormerrors.BadDecodeEntry.Wrapf("unexpected schema %d", schemaKey)
			}
		}

		if id > math.MaxUint32 {
			return nil, ormerrors.UnexpectedDecodePrefix.Wrapf("uint32 varint id out of range %d", id)
		}

		table, ok := m.tablesById[uint32(id)]
		if !ok {
			return nil, ormerrors.UnexpectedDecodePrefix.Wrapf("can't find table with id %d", id)
		}

		return table.DecodeKV(k, v)
	} else {
		return nil, ormerrors.UnexpectedDecodePrefix
	}
}

func (m ModuleSchema) EncodeKV(entry ormkv.Entry) (k, v []byte, err error) {
	tableName := entry.GetTableName()
	if tableName != "" {
		table, ok := m.tablesByName[tableName]
		if !ok {
			return nil, nil, ormerrors.BadDecodeEntry.Wrapf("can't find table %s", tableName)
		}

		return table.EncodeKV(entry)
	}

	if _, ok := entry.(ormkv.SchemaEntry); ok {
		return m.schemaKeyCodec.EncodeKV(entry)
	}

	return nil, nil, ormerrors.BadDecodeEntry.Wrapf("%s", entry)
}

var _ ormkv.Codec = &ModuleSchema{}
