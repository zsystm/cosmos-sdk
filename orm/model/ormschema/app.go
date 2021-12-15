package ormschema

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/orm/encoding/ormkv"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type AppSchema struct {
	moduleSchemas map[string]*ModuleSchema
	tablesByName  map[protoreflect.FullName]moduleTableEntry
}

type moduleTableEntry struct {
	module string
	table  ormtable.Table
	schema *ModuleSchema
}

func NewAppSchema(moduleSchemas map[string]*ModuleSchema) *AppSchema {
	schema := &AppSchema{
		moduleSchemas: moduleSchemas,
		tablesByName:  map[protoreflect.FullName]moduleTableEntry{},
	}

	for module, moduleSchema := range moduleSchemas {
		for name, table := range moduleSchema.tablesByName {
			schema.tablesByName[name] = moduleTableEntry{
				module: module,
				table:  table,
				schema: moduleSchema,
			}
		}
	}

	return schema
}

func (a AppSchema) DecodeEntry(module string, k, v []byte) (ormkv.Entry, error) {
	moduleSchema, ok := a.moduleSchemas[module]
	if !ok {
		return nil, fmt.Errorf("can't find module %s", module)
	}

	return moduleSchema.DecodeEntry(k, v)
}

func (a AppSchema) EncodeEntry(entry ormkv.Entry) (module string, k, v []byte, err error) {
	tableName := entry.GetTableName()
	tableEntry, ok := a.tablesByName[tableName]
	if !ok {
		return "", nil, nil, ormerrors.BadDecodeEntry.Wrapf("can't find table %s", tableName)
	}

	k, v, err = tableEntry.schema.EncodeEntry(entry)
	return tableEntry.module, k, v, err
}
