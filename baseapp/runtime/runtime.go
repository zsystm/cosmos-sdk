package runtime

import (
	"encoding/json"
	"fmt"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

type App struct {
	*baseapp.BaseApp
	beginBlockers []func(sdk.Context, abci.RequestBeginBlock)
	endBlockers   []func(sdk.Context, abci.RequestEndBlock) []abci.ValidatorUpdate
}

func (a App) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	ctx = ctx.WithEventManager(sdk.NewEventManager())

	for _, beginBlocker := range a.beginBlockers {
		beginBlocker(ctx, req)
	}

	return abci.ResponseBeginBlock{
		Events: ctx.EventManager().ABCIEvents(),
	}
}

func (a App) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	ctx = ctx.WithEventManager(sdk.NewEventManager())
	validatorUpdates := []abci.ValidatorUpdate{}

	for _, endBlocker := range a.endBlockers {
		moduleValUpdates := endBlocker(ctx, req)

		// use these validator updates if provided, the module manager assumes
		// only one module will update the validator set
		if len(moduleValUpdates) > 0 {
			if len(validatorUpdates) > 0 {
				panic("validator EndBlock updates already set by a previous module")
			}

			validatorUpdates = moduleValUpdates
		}
	}

	return abci.ResponseEndBlock{
		ValidatorUpdates: validatorUpdates,
		Events:           ctx.EventManager().ABCIEvents(),
	}
}

func (a App) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState map[string]json.RawMessage
	if err := json.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}
	// TODO: app.UpgradeKeeper.SetModuleVersionMap(ctx, app.mm.GetVersionMap())

	var validatorUpdates []abci.ValidatorUpdate
	ctx.Logger().Info("initializing blockchain state from genesis.json")
	for _, moduleName := range m.OrderInitGenesis {
		if genesisData[moduleName] == nil {
			continue
		}
		ctx.Logger().Debug("running initialization for module", "module", moduleName)

		moduleValUpdates := m.Modules[moduleName].InitGenesis(ctx, cdc, genesisData[moduleName])

		// use these validator updates if provided, the module manager assumes
		// only one module will update the validator set
		if len(moduleValUpdates) > 0 {
			if len(validatorUpdates) > 0 {
				panic("validator InitGenesis updates already set by a previous module")
			}
			validatorUpdates = moduleValUpdates
		}
	}

	// a chain must initialize with a non-empty validator set
	if len(validatorUpdates) == 0 {
		panic(fmt.Sprintf("validator set is empty after InitGenesis, please ensure at least one validator is initialized with a delegation greater than or equal to the DefaultPowerReduction (%d)", sdk.DefaultPowerReduction))
	}

	return abci.ResponseInitChain{
		Validators: validatorUpdates,
	}
}

func (a App) ExportAppStateAndValidators(forZeroHeight bool, jailAllowedAddrs []string) (servertypes.ExportedApp, error) {
	//TODO implement me
	panic("implement me")
}

func (a App) SimulationManager() *module.SimulationManager {
	//TODO implement me
	panic("implement me")
}

var _ SimappLikeApp = &App{}

func (a App) RegisterAPIRoutes(server *api.Server, config config.APIConfig) {}

func (a App) RegisterTxService(clientCtx client.Context) {}

func (a App) RegisterTendermintService(clientCtx client.Context) {}

var _ servertypes.Application = &App{}

type SimappLikeApp interface {
	// Exports the state of the application for a genesis file.
	ExportAppStateAndValidators(
		forZeroHeight bool, jailAllowedAddrs []string,
	) (servertypes.ExportedApp, error)

	// Helper for the simulation framework.
	SimulationManager() *module.SimulationManager
}
