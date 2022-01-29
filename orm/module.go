package orm

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"

	"github.com/cosmos/cosmos-sdk/orm/model/kv"

	modulev1alpha1 "github.com/cosmos/cosmos-sdk/api/cosmos/orm/module/v1alpha1"
	"github.com/cosmos/cosmos-sdk/container"
	"github.com/cosmos/cosmos-sdk/container/module"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
)

func init() {
	module.Register(&modulev1alpha1.Module{},
		module.Provide(func() {}),
	)
}

type ModuleSchema ormdb.ModuleSchema

func (m ModuleSchema) IsOnePerScopeType() {}
func (m ModuleSchema) IsModuleParamType() {}

var _ module.ParamType = ModuleSchema{}

func DefineModuleSchema(schema ModuleSchema) module.Option {
	return module.DefineParam(schema)
}

type needs struct {
	container.In

	ModuleSchemas map[container.Scope]ModuleSchema

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

func provideModuleDB(scope container.Scope, inputs needs) (provides, error) {
	schema, ok := inputs.ModuleSchemas[scope]
	if !ok {
		return provides{}, fmt.Errorf("missing module schema for module %s", scope.Name())
	}

	db, err := ormdb.NewModuleDB(ormdb.ModuleSchema(schema), ormdb.ModuleDBOptions{
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
