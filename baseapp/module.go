package baseapp

import (
	"github.com/cosmos/cosmos-sdk/container"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	store "github.com/cosmos/cosmos-sdk/store/types"
)

type Inputs struct {
	container.In

	Options []BaseAppOption
}

type Outputs struct {
	container.Out
	InterfaceRegistry codectypes.InterfaceRegistry
	Codec             codec.Codec
}

var Module = container.Options(
	container.Provide(
		provide,
		provideKVStoreKey,
		provideTransientStoreKey,
		provideMemoryStoreKey,
	),
)

type BaseAppOption struct{ F func(*BaseApp) }

func (BaseAppOption) IsAutoGroupType() {}

func provide(inputs Inputs) (Outputs, error) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	codec := codec.NewProtoCodec(interfaceRegistry)
	return Outputs{
		InterfaceRegistry: interfaceRegistry,
		Codec:             codec,
	}, nil
}

func provideKVStoreKey(key container.ModuleKey) (*store.KVStoreKey, BaseAppOption) {
	storeKey := store.NewKVStoreKey(key.Name())
	opt := func(app *BaseApp) {
		app.MountStores(storeKey)
	}
	return storeKey, BaseAppOption{opt}
}

func provideTransientStoreKey(key container.ModuleKey) (*store.TransientStoreKey, BaseAppOption) {
	storeKey := store.NewTransientStoreKey(key.Name())
	opt := func(app *BaseApp) {
		app.MountStores(storeKey)
	}
	return storeKey, BaseAppOption{opt}
}

func provideMemoryStoreKey(key container.ModuleKey) (*store.MemoryStoreKey, BaseAppOption) {
	storeKey := store.NewMemoryStoreKey(key.Name())
	opt := func(app *BaseApp) {
		app.MountStores(storeKey)
	}
	return storeKey, BaseAppOption{opt}
}
