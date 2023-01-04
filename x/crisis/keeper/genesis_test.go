package keeper_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistestutil "github.com/cosmos/cosmos-sdk/x/crisis/testutil"
	"github.com/cosmos/cosmos-sdk/x/crisis/types"
)

type genesisFixture struct {
	sdkCtx sdk.Context
	keeper keeper.Keeper
	cdc    codec.BinaryCodec
}

func initGenesisFixture(t assert.TestingT) *genesisFixture {
	f := &genesisFixture{}

	key := sdk.NewKVStoreKey(types.StoreKey)
	testCtx := testutil.DefaultContextWithDB(&testing.T{}, key, sdk.NewTransientStoreKey("transient_test"))
	encCfg := moduletestutil.MakeTestEncodingConfig(crisis.AppModuleBasic{})

	// gomock initializations
	ctrl := gomock.NewController(&testing.T{})
	f.cdc = codec.NewProtoCodec(encCfg.InterfaceRegistry)
	f.sdkCtx = testCtx.Ctx

	supplyKeeper := crisistestutil.NewMockSupplyKeeper(ctrl)
	f.keeper = *keeper.NewKeeper(f.cdc, key, 5, supplyKeeper, "", "")

	return f
}

func TestImportExportGenesis(t *testing.T) {
	t.Parallel()
	f := initGenesisFixture(t)

	// default params
	constantFee := sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000))
	err := f.keeper.SetConstantFee(f.sdkCtx, constantFee)
	assert.NilError(t, err)
	genesis := f.keeper.ExportGenesis(f.sdkCtx)

	// set constant fee to zero
	constantFee = sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(0))
	err = f.keeper.SetConstantFee(f.sdkCtx, constantFee)
	assert.NilError(t, err)

	f.keeper.InitGenesis(f.sdkCtx, genesis)
	newGenesis := f.keeper.ExportGenesis(f.sdkCtx)
	assert.DeepEqual(t, genesis, newGenesis)
}

func TestInitGenesis(t *testing.T) {
	t.Parallel()
	f := initGenesisFixture(t)

	genesisState := types.DefaultGenesisState()
	genesisState.ConstantFee = sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000))
	f.keeper.InitGenesis(f.sdkCtx, genesisState)

	constantFee := f.keeper.GetConstantFee(f.sdkCtx)
	assert.DeepEqual(t, genesisState.ConstantFee, constantFee)
}
