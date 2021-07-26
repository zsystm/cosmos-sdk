package app

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/container"
	"github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var StoreKeyProvider = container.Provide(
	func(scope container.Scope) (*types.KVStoreKey, func(*baseapp.BaseApp)) {
		key := types.NewKVStoreKey(scope.Name())
		return key, func(app *baseapp.BaseApp) {
			app.MountKVStores(map[string]*types.KVStoreKey{scope.Name(): key})
		}
	},
	func(scope container.Scope) (*types.KVStoreKey, func(*baseapp.BaseApp)) {
		key := types.NewKVStoreKey(scope.Name())
		return key, func(app *baseapp.BaseApp) {
			app.MountKVStores(map[string]*types.KVStoreKey{scope.Name(): key})
		}
	},
	func(scope container.Scope) (*types.MemoryStoreKey, func(*baseapp.BaseApp)) {
		name := fmt.Sprintf("mem:%s", scope.Name())
		key := types.NewMemoryStoreKey(name)
		return key, func(app *baseapp.BaseApp) {
			app.MountStore(key, sdk.StoreTypeMemory)
		}
	},
	func(scope container.Scope) (*types.TransientStoreKey, func(*baseapp.BaseApp)) {
		name := fmt.Sprintf("transient:%s", scope.Name())
		key := types.NewTransientStoreKey(name)
		return key, func(app *baseapp.BaseApp) {
			app.MountStore(key, sdk.StoreTypeTransient)
		}
	},
)
