package ormdb

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"math"
	"sort"

	"google.golang.org/protobuf/reflect/protodesc"

	"github.com/cosmos/cosmos-sdk/orm/encoding/encodeutil"

	"google.golang.org/protobuf/proto"

	ormv1alpha1 "github.com/cosmos/cosmos-sdk/api/cosmos/orm/v1alpha1"

	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/orm/encoding/ormkv"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
)

type ModuleSchema struct {
	FileDescriptors map[uint32]protoreflect.FileDescriptor
	Prefix          []byte
}

type ModuleDB struct {
	prefix         []byte
	filesById      map[uint32]*FileDescriptorDB
	tablesByName   map[protoreflect.FullName]ormtable.Table
	schemaSubspace *FileDescriptorDB
}

const (
	schemaSubspaceId uint32 = 0
)

type ModuleDBOptions struct {
	// TypeResolver is an optional type resolver to be used when unmarshaling
	// protobuf messages.
	TypeResolver ormtable.TypeResolver

	FileResolver protodesc.Resolver

	// JSONValidator is an optional validator that can be used for validating
	// messaging when using ValidateJSON. If it is nil, DefaultJSONValidator
	// will be used
	JSONValidator func(proto.Message) error

	GetBackend func(context.Context) (ormtable.Backend, error)

	GetReadBackend func(context.Context) (ormtable.ReadBackend, error)
}

func NewModuleDB(desc ModuleSchema, options ModuleDBOptions) (*ModuleDB, error) {
	prefix := desc.Prefix

	// the schema subspace is a private part of the store used for storing
	// important schema information for migrations and introspection
	schemaPrefix := encodeutil.AppendVarUInt32(prefix, schemaSubspaceId)
	schemaSubspace, err := NewFileDescriptorSchema(ormv1alpha1.File_cosmos_orm_v1alpha1_schema_proto, FileDescriptorDBOptions{
		Prefix:       schemaPrefix,
		ID:           1,
		TypeResolver: options.TypeResolver,
	})
	if err != nil {
		return nil, err
	}

	schema := &ModuleDB{
		prefix:         prefix,
		filesById:      map[uint32]*FileDescriptorDB{},
		tablesByName:   map[protoreflect.FullName]ormtable.Table{},
		schemaSubspace: schemaSubspace,
	}

	for id, fileDescriptor := range desc.FileDescriptors {
		opts := FileDescriptorDBOptions{
			ID:             id,
			Prefix:         prefix,
			TypeResolver:   options.TypeResolver,
			JSONValidator:  options.JSONValidator,
			GetBackend:     options.GetBackend,
			GetReadBackend: options.GetReadBackend,
		}

		if options.FileResolver != nil {
			// if a FileResolver is provided, we use that to resolve the file
			// and not the one provided as a different pinned file descriptor
			// may have been provided
			fileDescriptor, err = options.FileResolver.FindFileByPath(fileDescriptor.Path())
			if err != nil {
				return nil, err
			}
		}

		fdSchema, err := NewFileDescriptorSchema(fileDescriptor, opts)
		if err != nil {
			return nil, err
		}

		schema.filesById[id] = fdSchema
		for name, table := range fdSchema.tablesByName {
			if _, ok := schema.tablesByName[name]; ok {
				return nil, ormerrors.UnexpectedError.Wrapf("duplicate table %s", name)
			}

			schema.tablesByName[name] = table
		}
	}

	return schema, nil
}

func (m ModuleDB) DecodeEntry(k, v []byte) (ormkv.Entry, error) {
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

func (m ModuleDB) EncodeEntry(entry ormkv.Entry) (k, v []byte, err error) {
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

func (m ModuleDB) AutoMigrate(ctx context.Context) error {
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
		found, err := moduleFileTable.Get(ctx, &existing, id)
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
		err = moduleFileTable.Save(ctx, &ormv1alpha1.ModuleFileTable{
			Id:       id,
			FileName: filePath,
		})
		if err != nil {
			return err
		}

		err = file.AutoMigrate(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m ModuleDB) GetTable(message proto.Message) (ormtable.Table, error) {
	tableName := message.ProtoReflect().Descriptor().FullName()
	table, ok := m.tablesByName[tableName]
	if !ok {
		return nil, fmt.Errorf("table %T not found", tableName)
	}

	return table, nil
}

var _ ormkv.EntryCodec = &ModuleDB{}
var _ DB = &ModuleDB{}
