package ormdb

import (
	appv1alpha1 "github.com/cosmos/cosmos-sdk/api/cosmos/app/v1alpha1"
	ormv1alpha1 "github.com/cosmos/cosmos-sdk/api/cosmos/orm/v1alpha1"
	"github.com/cosmos/cosmos-sdk/container"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/types/kv"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
)

// AppDBOptions are options for constructing an AppDB.
type AppDBOptions struct {

	// TypeResolver is an optional type resolver to be used when unmarshaling
	// protobuf messages. If it is nil, protoregistry.GlobalTypes will be used.
	TypeResolver ormtable.TypeResolver

	// FileResolver is an optional file resolver that can be used to retrieve
	// pinned file descriptors that may be different from those available at
	// runtime. The file descriptor versions returned by this resolver will be
	// used instead of the ones provided at runtime by the ModuleSchema.
	FileResolver protodesc.Resolver

	// JSONValidator is an optional validator that can be used for validating
	// messaging when using ValidateJSON. If it is nil, DefaultJSONValidator
	// will be used
	JSONValidator func(proto.Message) error

	MultiStore MultiStore
}

type MultiStore interface {
	KVStore(module container.ModuleKey, storageType ormv1alpha1.StorageType) (kv.Store, error)
	ReadonlyKVStore(module string, storageType ormv1alpha1.StorageType) (kv.ReadonlyStore, error)
}

func NewAppDB(appConfig *appv1alpha1.Config, options AppDBOptions) (ormtable.Schema, error) {
	panic("TODO")
}
