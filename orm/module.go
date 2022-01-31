package orm

import (
	"embed"
	"fmt"

	modulev1alpha1 "github.com/cosmos/cosmos-sdk/api/cosmos/orm/module/v1alpha1"
	"github.com/cosmos/cosmos-sdk/app/module"
	"github.com/cosmos/cosmos-sdk/container"
	"github.com/cosmos/cosmos-sdk/orm/model/kv"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
)

//go:embed proto_image.bin.gz
var pinnedProtoImage embed.FS

func init() {
	module.Register(
		&modulev1alpha1.Module{},
		pinnedProtoImage,
		module.Provide(provideModuleDB),
	)
}

type needs struct {
	container.In

	ModuleSchemas map[container.ModuleKey]ormdb.ModuleSchema

	KVStore        KVStore
	IndexStore     IndexStore `optional:"true"`
	MemoryStore    TransientStore
	TransientStore TransientStore
	Hooks          ormtable.Hooks `optional:"true"`
}

type provides struct {
	container.Out

	DB ormdb.ModuleDB
}

func provideModuleDB(scope container.ModuleKey, inputs needs) (provides, error) {
	schema, ok := inputs.ModuleSchemas[scope]
	if !ok {
		return provides{}, fmt.Errorf("missing module schema for module %s", scope.Name())
	}

	db, err := ormdb.NewModuleDB(schema, ormdb.ModuleDBOptions{
		TypeResolver:   nil,
		FileResolver:   nil,
		JSONValidator:  nil,
		GetBackend:     nil,
		GetReadBackend: nil,
	})
	if err != nil {
		return provides{}, err
	}

	return provides{
		DB: db,
	}, nil
}

type KVStore kv.Store
type IndexStore kv.Store
type MemoryStore kv.Store
type TransientStore kv.Store
