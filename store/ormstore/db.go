package ormstore

import (
	"context"

	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/orm"
	"github.com/cosmos/cosmos-sdk/orm/model/ormschema"
	"github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type StoreKeyDB struct {
	key    *types.KVStoreKey
	schema ormschema.Schema
}

func NewStoreKeyDB(key *types.KVStoreKey, prefix []byte, fileDescriptors []protoreflect.FileDescriptor) (*StoreKeyDB, error) {
	schema, err := ormschema.NewModuleSchema(fileDescriptors, ormschema.ModuleSchemaOptions{
		Prefix: prefix,
	})
	return &StoreKeyDB{key: key, schema: schema}, err
}

func (s StoreKeyDB) OpenRead(ctx context.Context) (*orm.ReadDBConnection, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(s.key)
	wrapper := &kvStoreBackend{
		store: store,
	}
	return &orm.ReadDBConnection{
		Schema:      s.schema,
		ReadBackend: wrapper,
	}, nil
}

func (s StoreKeyDB) Open(ctx context.Context) (*orm.DBConnection, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(s.key)
	wrapper := &kvStoreBackend{
		store: store,
	}
	return &orm.DBConnection{
		ReadDBConnection: &orm.ReadDBConnection{
			Schema:      s.schema,
			ReadBackend: wrapper,
		},
		Backend: wrapper,
	}, nil
}

var _ orm.DB = StoreKeyDB{}
