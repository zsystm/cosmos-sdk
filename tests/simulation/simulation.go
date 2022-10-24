package simulation

import (
	"github.com/berachain/berachain-node/app/modules"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// BeraApp must conform to the simapp.App interface
var _ simapp.App = (*TestableSimApp)(nil)

// Name returns the name of the App.
func (app *TestableSimApp) Name() string { return app.BaseApp.Name() }

// Version returns the version of the App.
func (app *TestableSimApp) Version() string { return app.BaseApp.Version() }

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *TestableSimApp) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range modules.AccountPermissions {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// SimulationManager implements the SimulationApp interface.
func (app *TestableSimApp) SimulationManager() *module.SimulationManager {
	return app.sm
}
