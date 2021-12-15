package ormschema

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"sort"

	"github.com/cosmos/cosmos-sdk/orm/encoding/encodeutil"

	"github.com/cosmos/cosmos-sdk/orm/model/kvstore"

	"google.golang.org/protobuf/proto"

	ormv1alpha1 "github.com/cosmos/cosmos-sdk/api/cosmos/orm/v1alpha1"

	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/orm/encoding/ormkv"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
)

type ModuleSchema struct {
	prefix         []byte
	filesById      map[uint32]*FileDescriptorSchema
	tablesByName   map[protoreflect.FullName]ormtable.Table
	schemaSubspace *FileDescriptorSchema
}

const (
	schemaSubspaceId uint32 = 0
)

type ModuleSchemaOptions struct {
	Prefix            []byte
	FileDescriptorIds map[string]uint32
	// TypeResolver is an optional type resolver to be used when unmarshaling
	// protobuf messages.
	TypeResolver ormtable.TypeResolver

	// JSONValidator is an optional validator that can be used for validating
	// messaging when using ValidateJSON. If it is nil, DefaultJSONValidator
	// will be used
	JSONValidator func(proto.Message) error
}

func NewModuleSchema(fileDescriptors []protoreflect.FileDescriptor, options ModuleSchemaOptions) (*ModuleSchema, error) {
	prefix := options.Prefix

	// the schema subspace is a private part of the store used for storing
	// important schema information for migrations and introspection
	schemaPrefix := encodeutil.AppendVarUInt32(prefix, schemaSubspaceId)
	schemaSubspace, err := NewFileDescriptorSchema(ormv1alpha1.File_cosmos_orm_v1alpha1_schema_proto, FileDescriptorSchemaOptions{
		Prefix:       schemaPrefix,
		ID:           1,
		TypeResolver: options.TypeResolver,
	})
	if err != nil {
		return nil, err
	}

	schema := &ModuleSchema{
		prefix:         prefix,
		filesById:      map[uint32]*FileDescriptorSchema{},
		tablesByName:   map[protoreflect.FullName]ormtable.Table{},
		schemaSubspace: schemaSubspace,
	}

	for _, fileDescriptor := range fileDescriptors {
		opts := FileDescriptorSchemaOptions{
			Prefix:        prefix,
			TypeResolver:  options.TypeResolver,
			JSONValidator: options.JSONValidator,
		}

		if options.FileDescriptorIds != nil {
			if id, ok := options.FileDescriptorIds[fileDescriptor.Path()]; ok {
				opts.ID = id
			}
		}

		fdSchema, err := NewFileDescriptorSchema(fileDescriptor, opts)
		if err != nil {
			return nil, err
		}

		schema.filesById[fdSchema.id] = fdSchema
		for name, table := range fdSchema.tablesByName {
			if _, ok := schema.tablesByName[name]; ok {
				return nil, ormerrors.UnexpectedError.Wrapf("duplicate table %s", name)
			}

			schema.tablesByName[name] = table
		}
	}

	return schema, nil
}

func (m ModuleSchema) DecodeEntry(k, v []byte) (ormkv.Entry, error) {
	r := bytes.NewReader(k)
	err := encodeutil.SkipPrefix(r, m.prefix)
	if err != nil {
		return nil, err
	}

	id, err := binary.ReadUvarint(r)
	if err != nil {
		return nil, err
	}

	if id > math.MaxUint32 {
		return nil, ormerrors.UnexpectedDecodePrefix.Wrapf("uint32 varint id out of range %d", id)
	}

	// schema sub-store
	if uint32(id) == schemaSubspaceId {
		return m.schemaSubspace.DecodeEntry(k, v)
	}

	fileSchema, ok := m.filesById[uint32(id)]
	if !ok {
		return nil, ormerrors.UnexpectedDecodePrefix.Wrapf("can't find FileDescriptor schema with id %d", id)
	}

	return fileSchema.DecodeEntry(k, v)
}

func (m ModuleSchema) EncodeEntry(entry ormkv.Entry) (k, v []byte, err error) {
	tableName := entry.GetTableName()
	table, ok := m.tablesByName[tableName]
	if !ok {
		table, ok = m.schemaSubspace.tablesByName[tableName]
		if !ok {
			return nil, nil, ormerrors.BadDecodeEntry.Wrapf("can't find table %s", tableName)
		}
	}

	return table.EncodeEntry(entry)
}

func (m ModuleSchema) AutoMigrate(store kvstore.Backend) error {
	moduleFileTable := m.schemaSubspace.GetTable(&ormv1alpha1.ModuleFileTable{})
	if moduleFileTable == nil {
		return ormerrors.UnexpectedError.Wrapf("missing ModuleFileTable")
	}

	var sortedIds []int
	for id := range m.filesById {
		sortedIds = append(sortedIds, int(id))
	}
	sort.Ints(sortedIds)

	for _, id := range sortedIds {
		id := uint32(id)
		file := m.filesById[id]

		var existing ormv1alpha1.ModuleFileTable
		found, err := moduleFileTable.Get(store, encodeutil.ValuesOf(id), &existing)
		if err != nil {
			return err
		}

		filePath := file.fileDescriptor.Path()
		if found {
			if existing.FileName != filePath {
				return ormerrors.MigrationError.Wrapf(
					"file descriptor %s with at ID %d already exists, can't replace with %s",
					existing.FileName,
					id,
					filePath,
				)
			}
		}

		// because of the unique index on file_name, this will fail
		// if the file was already registered with a different id
		err = moduleFileTable.Save(store, &ormv1alpha1.ModuleFileTable{
			Id:       id,
			FileName: filePath,
		}, ormtable.SAVE_MODE_INSERT)
		if err != nil {
			return err
		}

		err = file.AutoMigrate(store)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m ModuleSchema) GetTable(message proto.Message) (ormtable.Table, error) {
	tableName := message.ProtoReflect().Descriptor().FullName()
	table, ok := m.tablesByName[tableName]
	if !ok {
		return nil, fmt.Errorf("table %T not found", tableName)
	}

	return table, nil
}

var _ ormkv.EntryCodec = &ModuleSchema{}
var _ Schema = &ModuleSchema{}
