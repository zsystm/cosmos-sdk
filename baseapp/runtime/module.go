package runtime

import (
	"fmt"
	"io"

	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	runtimev1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/runtime/v1"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/container"
	coremodule "github.com/cosmos/cosmos-sdk/core/module"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
)

type BaseAppOption func(*baseapp.BaseApp)

func (b BaseAppOption) IsAutoGroupType() {}

type appBuilder struct {
	storeKeys         []storetypes.StoreKey
	interfaceRegistry codectypes.InterfaceRegistry
	cdc               codec.Codec
	amino             *codec.LegacyAmino
}

func (a *appBuilder) registerStoreKey(key storetypes.StoreKey) {
	a.storeKeys = append(a.storeKeys, key)
}

func init() {
	coremodule.Register(&runtimev1.Module{},
		coremodule.Provide(
			provideBuilder,
			provideApp,
			provideKVStoreKey,
			provideTransientStoreKey,
			provideMemoryStoreKey,
		),
	)
}

func provideBuilder(moduleBasics map[string]module.AppModuleBasicWiringWrapper) (
	codectypes.InterfaceRegistry,
	codec.Codec,
	*codec.LegacyAmino,
	*appBuilder) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	amino := codec.NewLegacyAmino()
	for _, wrapper := range moduleBasics {
		wrapper.RegisterInterfaces(interfaceRegistry)
		wrapper.RegisterLegacyAminoCodec(amino)
	}
	cdc := codec.NewProtoCodec(interfaceRegistry)
	builder := &appBuilder{
		storeKeys:         nil,
		interfaceRegistry: interfaceRegistry,
		cdc:               cdc,
		amino:             amino,
	}

	return interfaceRegistry, cdc, amino, builder
}

type AppCreator struct {
	app *App
}

func (a *AppCreator) RegisterModules(modules ...module.AppModule) error {
	for _, appModule := range modules {
		if _, ok := a.app.mm.Modules[appModule.Name()]; ok {
			return fmt.Errorf("module named %q already exists", appModule.Name())
		}
		a.app.mm.Modules[appModule.Name()] = appModule
	}
	return nil
}

func (a *AppCreator) Create(logger log.Logger, db dbm.DB, traceStore io.Writer, baseAppOptions ...func(*baseapp.BaseApp)) *App {
	for _, option := range a.app.baseAppOptions {
		baseAppOptions = append(baseAppOptions, option)
	}
	bApp := baseapp.NewBaseApp(a.app.config.AppName, logger, db, baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(a.app.builder.interfaceRegistry)
	bApp.MountStores(a.app.builder.storeKeys...)

	a.app.BaseApp = bApp
	return a.app
}

func (a *AppCreator) Finish(loadLatest bool) error {
	if a.app == nil {
		return fmt.Errorf("app not created yet, can't finish")
	}

	a.app.mm.SetOrderInitGenesis(a.app.config.InitGenesis...)
	a.app.mm.SetOrderBeginBlockers(a.app.config.BeginBlockers...)
	a.app.mm.SetOrderEndBlockers(a.app.config.EndBlockers...)
	a.app.SetBeginBlocker(a.app.mm.BeginBlock)
	a.app.SetEndBlocker(a.app.mm.EndBlock)
	a.app.SetInitChainer(a.app.InitChainer)

	if loadLatest {
		if err := a.app.LoadLatestVersion(); err != nil {
			return err
		}
	}

	return nil
}

func provideApp(config *runtimev1.Module, builder *appBuilder, modules map[string]module.AppModuleWiringWrapper, baseAppOptions []BaseAppOption) *AppCreator {
	mm := &module.Manager{Modules: map[string]module.AppModule{}}
	for name, wrapper := range modules {
		mm.Modules[name] = wrapper.AppModule
	}
	return &AppCreator{
		app: &App{
			BaseApp:        nil,
			baseAppOptions: baseAppOptions,
			config:         config,
			builder:        builder,
			mm:             mm,
			beginBlockers:  nil,
			endBlockers:    nil,
		},
	}
}

func provideKVStoreKey(key container.ModuleKey, builder *appBuilder) *storetypes.KVStoreKey {
	storeKey := storetypes.NewKVStoreKey(key.Name())
	builder.registerStoreKey(storeKey)
	return storeKey
}

func provideTransientStoreKey(key container.ModuleKey, builder *appBuilder) *storetypes.TransientStoreKey {
	storeKey := storetypes.NewTransientStoreKey(fmt.Sprintf("transient:%s", key.Name()))
	builder.registerStoreKey(storeKey)
	return storeKey
}

func provideMemoryStoreKey(key container.ModuleKey, builder *appBuilder) *storetypes.MemoryStoreKey {
	storeKey := storetypes.NewMemoryStoreKey(fmt.Sprintf("memory:%s", key.Name()))
	builder.registerStoreKey(storeKey)
	return storeKey
}
