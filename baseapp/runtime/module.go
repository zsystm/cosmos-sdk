package runtime

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/spf13/cast"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	runtimev1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/runtime/v1"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/container"
	coremodule "github.com/cosmos/cosmos-sdk/core/module"
	"github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/snapshots"
	snapshottypes "github.com/cosmos/cosmos-sdk/snapshots/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
)

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

var Module = container.Options(
	container.Provide(),
)

func provideBuilder() (
	codectypes.InterfaceRegistry,
	codec.Codec,
	*codec.LegacyAmino,
	*appBuilder) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)
	amino := codec.NewLegacyAmino()
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

func (a *AppCreator) Create(logger log.Logger, db dbm.DB, traceStore io.Writer, appOpts servertypes.AppOptions, baseAppOptions ...func(*baseapp.BaseApp)) *App {
	var cache sdk.MultiStorePersistentCache

	if cast.ToBool(appOpts.Get(server.FlagInterBlockCache)) {
		cache = store.NewCommitKVStoreCacheManager()
	}

	skipUpgradeHeights := make(map[int64]bool)
	for _, h := range cast.ToIntSlice(appOpts.Get(server.FlagUnsafeSkipUpgrades)) {
		skipUpgradeHeights[int64(h)] = true
	}

	pruningOpts, err := server.GetPruningOptionsFromFlags(appOpts)
	if err != nil {
		panic(err)
	}

	snapshotDir := filepath.Join(cast.ToString(appOpts.Get(flags.FlagHome)), "data", "snapshots")
	snapshotDB, err := dbm.NewDB("metadata", server.GetAppDBBackend(appOpts), snapshotDir)
	if err != nil {
		panic(err)
	}
	snapshotStore, err := snapshots.NewStore(snapshotDB, snapshotDir)
	if err != nil {
		panic(err)
	}

	snapshotOptions := snapshottypes.NewSnapshotOptions(
		cast.ToUint64(appOpts.Get(server.FlagStateSyncSnapshotInterval)),
		cast.ToUint32(appOpts.Get(server.FlagStateSyncSnapshotKeepRecent)),
	)

	baseAppOptions = append(baseAppOptions,
		baseapp.SetPruning(pruningOpts),
		baseapp.SetMinGasPrices(cast.ToString(appOpts.Get(server.FlagMinGasPrices))),
		baseapp.SetHaltHeight(cast.ToUint64(appOpts.Get(server.FlagHaltHeight))),
		baseapp.SetHaltTime(cast.ToUint64(appOpts.Get(server.FlagHaltTime))),
		baseapp.SetMinRetainBlocks(cast.ToUint64(appOpts.Get(server.FlagMinRetainBlocks))),
		baseapp.SetInterBlockCache(cache),
		baseapp.SetTrace(cast.ToBool(appOpts.Get(server.FlagTrace))),
		baseapp.SetIndexEvents(cast.ToStringSlice(appOpts.Get(server.FlagIndexEvents))),
		baseapp.SetSnapshot(snapshotStore, snapshotOptions),
	)

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

func provideApp(config *runtimev1.Module, builder *appBuilder, modules map[string]module.AppModuleWiringWrapper) *AppCreator {
	mm := &module.Manager{}
	for name, wrapper := range modules {
		mm.Modules[name] = wrapper.AppModule
	}
	return &AppCreator{
		app: &App{
			BaseApp:       nil,
			config:        config,
			builder:       builder,
			mm:            mm,
			beginBlockers: nil,
			endBlockers:   nil,
		},
	}
}

func provideKVStoreKey(key container.ModuleKey, builder *appBuilder) *storetypes.KVStoreKey {
	storeKey := storetypes.NewKVStoreKey(key.Name())
	builder.registerStoreKey(storeKey)
	return storeKey
}

func provideTransientStoreKey(key container.ModuleKey, builder *appBuilder) *storetypes.TransientStoreKey {
	storeKey := storetypes.NewTransientStoreKey(key.Name())
	builder.registerStoreKey(storeKey)
	return storeKey
}

func provideMemoryStoreKey(key container.ModuleKey, builder *appBuilder) *storetypes.MemoryStoreKey {
	storeKey := storetypes.NewMemoryStoreKey(key.Name())
	builder.registerStoreKey(storeKey)
	return storeKey
}
