package ormstore

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm"
	"github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type StoreKeyDB struct {
	key    *types.KVStoreKey
	schema orm.Schema
}

func (s StoreKeyDB) OpenRead(ctx context.Context) (*orm.ReadClient, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(s.key)
	wrapper := &kvStoreStore{
		store: store,
	}
	return &orm.ReadClient{
		Schema:      s.schema,
		ReadBackend: wrapper,
	}, nil
}

func (s StoreKeyDB) Open(ctx context.Context) (*orm.Client, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(s.key)
	wrapper := &kvStoreStore{
		store: store,
	}
	return &orm.Client{
		ReadClient: &orm.ReadClient{
			Schema:      s.schema,
			ReadBackend: wrapper,
		},
		Backend: wrapper,
	}, nil
}

var _ orm.DB = StoreKeyDB{}
