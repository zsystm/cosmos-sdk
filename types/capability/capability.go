package capability

import (
	"context"
	"cosmossdk.io/core/blockinfo"
	"cosmossdk.io/core/event"
	"cosmossdk.io/core/gas"
	"cosmossdk.io/core/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"reflect"
)

type Context interface {
	ModuleCapability() ModuleCapability
	WithModuleCapability(capabilities ModuleCapability) Context
}

const SdkContextKey string = "sdk-context"

// Module Context

type ModuleCapability map[string]bool

type ContextFactory[T Context] struct {
	Make func(ctx context.Context) T
}

func NewContextFactory[T Context]() ContextFactory[T] {
	return ContextFactory[T]{
		Make: func(ctx context.Context) T {
			if sdkCtx, ok := ctx.(T); ok && sdkCtx.ModuleCapability() != nil {
				return sdkCtx
			}
			sdkCtx := ctx.Value(SdkContextKey)
			if sdkCtx == nil {
				panic("sdk context not found")
			}
			c := sdkCtx.(T)
			t := reflect.TypeOf((*T)(nil)).Elem()
			capabilities := make(ModuleCapability)
			for i := 0; i < t.NumMethod(); i++ {
				name := t.Method(i).Name
				capabilities[name] = true
			}
			moduleContext := c.WithModuleCapability(capabilities)
			return moduleContext.(T)
		}}
}

const BlockInfoCapabilityKey string = "BlockInfoService"

type BlockInfoService interface {
	BlockInfoService() blockinfo.Service
}

const KVStoreCapabilityKey string = "KVStoreService"

type KVStoreService interface {
	KVStoreService(key storetypes.KVStoreKey) store.KVStoreService
}

const EventCapabilityKey string = "EventService"

type EventService interface {
	EventService() event.Service
}

const GasCapabilityKey string = "GasService"

type GasService interface {
	GasService() gas.Service
}
