package slashing_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/depinject"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/testutil/mock"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	"github.com/cosmos/cosmos-sdk/x/slashing/testutil"
	"github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

var (
	priv1 = secp256k1.GenPrivKey()
	addr1 = sdk.AccAddress(priv1.PubKey().Address())

	valKey  = ed25519.GenPrivKey()
	valAddr = sdk.AccAddress(valKey.PubKey().Address())
)

func checkValidator(t *testing.T, ctxCheck sdk.Context, stakingKeeper *stakingkeeper.Keeper, addr sdk.AccAddress, expFound bool) stakingtypes.Validator {
	validator, found := stakingKeeper.GetValidator(ctxCheck, sdk.ValAddress(addr))
	require.Equal(t, expFound, found)
	return validator
}

func checkValidatorSigningInfo(t *testing.T, ctxCheck sdk.Context, slashingKeeper slashingkeeper.Keeper, addr sdk.ConsAddress, expFound bool) types.ValidatorSigningInfo {
	signingInfo, found := slashingKeeper.GetValidatorSigningInfo(ctxCheck, addr)
	require.Equal(t, expFound, found)
	return signingInfo
}

// CheckBalance checks the balance of an account.
func CheckBalance(t *testing.T, ctxCheck sdk.Context, bankKeeper bankkeeper.Keeper, addr sdk.AccAddress, balances sdk.Coins) {
	require.True(t, balances.IsEqual(bankKeeper.GetAllBalances(ctxCheck, addr)))
}

func TestSlashingMsgs(t *testing.T) {
	priv1 = mock.NewPV()
	// pubKey, err := priv1.GetPubKey(context.TODO())
	// require.NoError(t, err)

	valKey = ed25519.GenPrivKey()
	valAddr = sdk.AccAddress(valKey.PubKey().Address())

	genTokens := sdk.TokensFromConsensusPower(42, sdk.DefaultPowerReduction)
	bondTokens := sdk.TokensFromConsensusPower(10, sdk.DefaultPowerReduction)
	genCoin := sdk.NewCoin(sdk.DefaultBondDenom, genTokens)
	bondCoin := sdk.NewCoin(sdk.DefaultBondDenom, bondTokens)

	acc1 := &authtypes.BaseAccount{
		Address: addr1.String(),
	}
	accs := authtypes.GenesisAccounts{acc1}
	balances := []banktypes.Balance{
		{
			Address: addr1.String(),
			Coins:   sdk.Coins{genCoin},
		},
	}

	var codec codec.Codec
	var appBuilder *runtime.AppBuilder
	err := depinject.Inject(testutil.AppConfig, &appBuilder, &codec)
	require.NoError(t, err)

	//
	// create genesis and validator
	//
	// privVal := mock.NewPV()
	// pubKey, err := privVal.GetPubKey(context.TODO())
	// require.NoError(t, err)

	pubKey, err := priv1.GetPubKey(context.TODO())
	// create validator set with single validator
	testValidator := tmtypes.NewValidator(pubKey, 1)
	valSet := tmtypes.NewValidatorSet([]*tmtypes.Validator{testValidator})

	genesisState, err := simtestutil.GenesisStateWithValSet(codec, appBuilder.DefaultGenesis(), valSet, accs, balances...)
	require.NoError(t, err)

	var bankKeeper bankkeeper.Keeper
	var stakingKeeper *stakingkeeper.Keeper
	var slashingKeeper slashingkeeper.Keeper
	app, err := simtestutil.SetupWithGenesisState(
		testutil.AppConfig,
		nil,
		valSet,
		genesisState,
		false,
		&bankKeeper,
		&stakingKeeper,
		&slashingKeeper,
	)
	require.NoError(t, err)
	ctxCheck := app.BaseApp.NewContext(true, tmproto.Header{})
	CheckBalance(t, ctxCheck, bankKeeper, addr1, sdk.Coins{genCoin})

	stakingKeeper.SetHooks(stakingtypes.NewMultiStakingHooks(slashingKeeper.Hooks()))

	description := stakingtypes.NewDescription("foo_moniker", "", "", "", "")
	commission := stakingtypes.NewCommissionRates(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec())

	createValidatorMsg, err := stakingtypes.NewMsgCreateValidator(
		sdk.ValAddress(addr1), valKey.PubKey(), bondCoin, description, commission, sdk.OneInt(),
	)
	require.NoError(t, err)

	header := tmproto.Header{Height: app.LastBlockHeight() + 1}
	txGen := simapp.MakeTestEncodingConfig().TxConfig
	_, _, err = simapp.SignCheckDeliver(t, txGen, app.BaseApp, header, []sdk.Msg{createValidatorMsg}, "", []uint64{0}, []uint64{0}, true, true, priv1)
	require.NoError(t, err)
	CheckBalance(t, ctxCheck, bankKeeper, addr1, sdk.Coins{genCoin.Sub(bondCoin)})

	header = tmproto.Header{Height: app.LastBlockHeight() + 1}
	app.BeginBlock(abci.RequestBeginBlock{Header: header})

	validator := checkValidator(t, ctxCheck, stakingKeeper, addr1, true)
	require.Equal(t, sdk.ValAddress(addr1).String(), validator.OperatorAddress)
	require.Equal(t, stakingtypes.Bonded, validator.Status)
	require.True(sdk.IntEq(t, bondTokens, validator.BondedTokens()))
	unjailMsg := &types.MsgUnjail{ValidatorAddr: sdk.ValAddress(addr1).String()}

	checkValidatorSigningInfo(t, ctxCheck, slashingKeeper, sdk.ConsAddress(valAddr), true)

	// unjail should fail with unknown validator
	header = tmproto.Header{Height: app.LastBlockHeight() + 1}
	_, res, err := simapp.SignCheckDeliver(t, txGen, app.BaseApp, header, []sdk.Msg{unjailMsg}, "", []uint64{0}, []uint64{1}, false, false, priv1)
	require.Error(t, err)
	require.Nil(t, res)
	require.True(t, errors.Is(types.ErrValidatorNotJailed, err))
}
