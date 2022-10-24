package simulation

import (
	"encoding/json"
	"fmt"
	"time"

	beraapp "github.com/berachain/berachain-node/app"
	bictypes "github.com/berachain/berachain-node/pkg/modules/bic/types"
	beradconfig "github.com/berachain/berachain-node/pkg/sdk/config"
	constants "github.com/berachain/berachain-node/pkg/sdk/constants"
	"github.com/berachain/berachain-node/pkg/testing"
	testinghelper "github.com/berachain/berachain-node/pkg/testing/utils"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	pruningtypes "github.com/cosmos/cosmos-sdk/pruning/types"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/evmos/ethermint/crypto/hd"
	evmenc "github.com/evmos/ethermint/encoding"
	"github.com/evmos/ethermint/testutil/network"
	ethermint "github.com/evmos/ethermint/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	appparams "github.com/cosmos/cosmos-sdk/simapp/params"

	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
)

var (
	ModuleBasics       = module.NewBasicManager(simapp.AppModuleBasics()...)
	TestEncodingConfig = evmenc.MakeConfig(ModuleBasics)
)

// TestableSimApp must conform to our custom testhelper.TestableApp interface
var _ testinghelper.TestableApp = (*TestableSimApp)(nil)

type TestableSimApp struct {
	*simapp.SimApp

	// simulation manager
	sm *module.SimulationManager
}

// GetCodec returns the application codec.
func (app *TestableSimApp) GetCodec() codec.Codec {
	return app.Codec
}

// NewDefaultsimapp.GenesisState generates the default state for the application.
func (app *TestableSimApp) GetDefaultGenesisState() simapp.GenesisState {
	return ModuleBasics.DefaultGenesis(app.Codec)
}

// GetModuleBasics returns the module basics for the application.
func (app *TestableSimApp) GetModuleBasics() *module.BasicManager {
	return &ModuleBasics
}

// GetTestEncodingConfig returns the test encoding config for the application.
func (app *TestableSimApp) GetTestEncodingConfig() simappparams.EncodingConfig {
	return TestEncodingConfig
}

// GetMintModuleAccount returns the "mint" module account
func (app *TestableSimApp) GetMintModuleAccount() string {
	return bictypes.ModuleName
}

// GetBankKeeper returns the bank keeper for this application
func (app *TestableSimApp) GetBankKeeper() bankkeeper.Keeper {
	return app.Keepers.BankKeeper
}

// LegacyAmino returns TestableBeraApp's amino codec.
func (app *TestableSimApp) LegacyAmino() *codec.LegacyAmino {
	return app.EncodingConfig.Amino
}

// Setup initializes a new BeraApp. A Nop logger is set in BeraApp.
func Setup(
	isCheckTx bool,
	chainID string,
	patchGenesis func(*TestableSimApp, simapp.GenesisState) simapp.GenesisState,
) *TestableSimApp {
	cfg := sdk.GetConfig()
	appparams.SetBech32Prefixes(cfg)
	appparams.SetBip44CoinType(cfg)

	return SetupWithDB(isCheckTx, chainID, patchGenesis, dbm.NewMemDB())
}

// SetupWithDB initializes a new BeraApp. A Nop logger is set in BeraApp.
func SetupWithDB(
	isCheckTx bool,
	chainID string,
	patchGenesis func(*TestableSimApp, simapp.GenesisState) simapp.GenesisState,
	db dbm.DB,
) *TestableSimApp {
	app := &TestableSimApp{
		beraapp.NewBeraApp(log.NewNopLogger(), db, nil, true, map[int64]bool{}, beradconfig.DefaultNodeHome, 5,
			evmenc.MakeConfig(ModuleBasics), simapp.EmptyAppOptions{}), nil,
	}

	// NOTE: this is not required apps that don't use the simulator for fuzz testing transactions
	overrideModules := map[string]module.AppModuleSimulation{
		authtypes.ModuleName: auth.NewAppModule(app.Codec, *app.AccountKeeper, authsims.RandomGenesisAccounts),
	}

	// Setup Simulation Manager for Sim Testing
	app.sm = module.NewSimulationManagerFromAppModules(app.Mm.Modules, overrideModules)

	// Setup Store Decoders
	app.sm.RegisterStoreDecoders()
	if !isCheckTx {
		// init chain must be called to stop deliverState from being nil
		genesisState := testinghelper.NewTestGenesisState(app)
		if patchGenesis != nil {
			genesisState = patchGenesis(app, genesisState)
		}

		stateBytes, err := json.MarshalIndent(genesisState, "", " ")
		if err != nil {
			panic(err)
		}

		// Initialize the chain
		app.InitChain(
			abci.RequestInitChain{
				ChainId:         chainID,
				Validators:      []abci.ValidatorUpdate{},
				ConsensusParams: testing.DefaultConsensusParams,
				AppStateBytes:   stateBytes,
			},
		)
	}

	return app
}

// NewAppConstructor returns a new simapp AppConstructor
func NewAppConstructor(encodingCfg simappparams.EncodingConfig) network.AppConstructor {
	return func(val network.Validator) servertypes.Application {
		return beraapp.NewBeraApp(
			val.Ctx.Logger, dbm.NewMemDB(), nil, true, make(map[int64]bool), val.Ctx.Config.RootDir, 0,
			encodingCfg,
			simapp.EmptyAppOptions{},
			baseapp.SetPruning(pruningtypes.NewPruningOptionsFromString(val.AppConfig.Pruning)),
			baseapp.SetMinGasPrices(val.AppConfig.MinGasPrices),
		)
	}
}

// DefaultConfig returns a sane default configuration suitable for nearly all
// testing requirements.
func DefaultConfig() network.Config {
	encCfg := evmenc.MakeConfig(ModuleBasics)

	return network.Config{
		Codec:             encCfg.Codec,
		TxConfig:          encCfg.TxConfig,
		LegacyAmino:       encCfg.Amino,
		InterfaceRegistry: encCfg.InterfaceRegistry,
		AccountRetriever:  authtypes.AccountRetriever{},
		AppConstructor:    NewAppConstructor(encCfg),
		GenesisState:      ModuleBasics.DefaultGenesis(encCfg.Codec),
		TimeoutCommit:     2 * time.Second,
		ChainID:           beradconfig.GenerateRandomChainID(),
		NumValidators:     1,
		BondDenom:         constants.AttoBeraUnit,
		MinGasPrices:      fmt.Sprintf("0.000006%s", constants.AttoBeraUnit),
		AccountTokens:     sdk.TokensFromConsensusPower(1000, ethermint.PowerReduction),
		StakingTokens:     sdk.TokensFromConsensusPower(500, ethermint.PowerReduction),
		BondedTokens:      sdk.TokensFromConsensusPower(100, ethermint.PowerReduction),
		PruningStrategy:   pruningtypes.PruningOptionNothing,
		CleanupDir:        true,
		SigningAlgo:       string(hd.EthSecp256k1Type),
		KeyringOptions:    []keyring.Option{hd.EthSecp256k1Option()},
		PrintMnemonic:     false,
	}
}
