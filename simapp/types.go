package simapp

import (
	"github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

// App implements the common methods for a Cosmos SDK-based application
// specific blockchain.
type App interface {
	// Exports the state of the application for a genesis file.
	ExportAppStateAndValidators(
		forZeroHeight bool, jailAllowedAddrs []string,
	) (types.ExportedApp, error)

	// Helper for the simulation framework.
	SimulationManager() *module.SimulationManager
}
