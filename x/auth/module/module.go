package module

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/app"
	"github.com/cosmos/cosmos-sdk/app/compat"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/container"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/simulation"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	_ app.TypeProvider = Module{}
)

type inputs struct {
	container.StructArgs

	Codec    codec.Codec
	Key      *sdk.KVStoreKey
	Subspace paramtypes.Subspace
}

type outputs struct {
	container.StructArgs

	Handler    app.Handler `group:"app"`
	ViewKeeper types.ViewKeeper
	Keeper     types.Keeper
}

type cliCommands struct {
	container.StructArgs

	TxCmd    *cobra.Command   `group:"tx"`
	QueryCmd []*cobra.Command `group:"query,flatten"`
}

func (m Module) RegisterTypes(registry codectypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

func (Module) provideAccountRetriever() client.AccountRetriever {
	return types.AccountRetriever{}
}

func (Module) provideCLICommands() cliCommands {
	am := auth.AppModuleBasic{}
	return cliCommands{
		TxCmd: am.GetTxCmd(),
		QueryCmd: []*cobra.Command{
			am.GetQueryCmd(),
			authcmd.GetAccountCmd(),
		},
	}
}

func (m Module) provideAppHandler(inputs inputs) (outputs, error) {
	var accCtr types.AccountConstructor
	if m.AccountConstructor != nil {
		err := inputs.Codec.UnpackAny(m.AccountConstructor, &accCtr)
		if err != nil {
			return outputs{}, err
		}
	} else {
		accCtr = DefaultAccountConstructor{}
	}

	perms := map[string][]string{}
	for _, perm := range m.Permissions {
		perms[perm.Address] = perm.Permissions
	}

	var randomGenesisAccountsProvider types.RandomGenesisAccountsProvider
	if m.RandomGenesisAccountsProvider != nil {
		err := inputs.Codec.UnpackAny(m.RandomGenesisAccountsProvider, &randomGenesisAccountsProvider)
		if err != nil {
			return outputs{}, err
		}
	} else {
		randomGenesisAccountsProvider = DefaultRandomGenesisAccountsProvider{}
	}

	keeper := authkeeper.NewAccountKeeper(
		inputs.Codec,
		inputs.Key,
		inputs.Subspace,
		func() types.AccountI {
			return accCtr.NewAccount()
		},
		perms,
	)
	appMod := auth.NewAppModule(inputs.Codec, keeper, func(simState *module.SimulationState) types.GenesisAccounts {
		return randomGenesisAccountsProvider.RandomGenesisAccounts(simState)
	})

	return outputs{
		ViewKeeper: viewOnlyKeeper{keeper},
		Keeper:     keeper,
		Handler:    compat.AppModuleHandler(appMod),
	}, nil
}

func (m DefaultAccountConstructor) NewAccount() types.AccountI {
	return &types.BaseAccount{}
}

func (m DefaultRandomGenesisAccountsProvider) RandomGenesisAccounts(simState *module.SimulationState) types.GenesisAccounts {
	return simulation.RandomGenesisAccounts(simState)
}
