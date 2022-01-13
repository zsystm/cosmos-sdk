package ormstore

import (
	"github.com/cosmos/cosmos-sdk/orm/model/ormschema"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

//type DB struct {
//	key    *types.KVStoreKey
//	schema ormschema.Schema
//}

func KVStoreDB(desc ormschema.ModuleDescriptor, key storetypes.StoreKey, resolver ormtable.TypeResolver) {

}
