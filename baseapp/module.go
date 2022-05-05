package baseapp

import (
	"github.com/cosmos/cosmos-sdk/container"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	store "github.com/cosmos/cosmos-sdk/store/types"
)

type Inputs struct {
	container.In
}

type Outputs struct {
	container.Out
	InterfaceRegistry codectypes.InterfaceRegistry
	Codec             codec.Codec
	StoreKeyRegistrar func(store.StoreKey)
}

var Module = container.Options(
	container.Provide(
		provide,
		provideKVStoreKey,
		provideTransientStoreKey,
		provideMemoryStoreKey,
	),
)

func provide(inputs Inputs) (Outputs, error) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	codec := codec.NewProtoCodec(interfaceRegistry)
	return Outputs{
		InterfaceRegistry: interfaceRegistry,
		Codec:             codec,
		StoreKeyRegistrar: func(store.StoreKey) {},
	}, nil
}

func provideKVStoreKey(key container.ModuleKey, registrar func(store.StoreKey)) *store.KVStoreKey {
	storeKey := store.NewKVStoreKey(key.Name())
	registrar(storeKey)
	return storeKey
}

func provideTransientStoreKey(key container.ModuleKey, registrar func(store.StoreKey)) *store.TransientStoreKey {
	storeKey := store.NewTransientStoreKey(key.Name())
	registrar(storeKey)
	return storeKey
}

func provideMemoryStoreKey(key container.ModuleKey, registrar func(store.StoreKey)) *store.MemoryStoreKey {
	storeKey := store.NewMemoryStoreKey(key.Name())
	registrar(storeKey)
	return storeKey
}
