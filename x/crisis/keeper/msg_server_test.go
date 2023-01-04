package keeper_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistestutil "github.com/cosmos/cosmos-sdk/x/crisis/testutil"
	"github.com/cosmos/cosmos-sdk/x/crisis/types"
)

type keeperFixture struct {
	ctx        sdk.Context
	authKeeper *crisistestutil.MockSupplyKeeper
	keeper     *keeper.Keeper
}

func initKeeperFixture(t assert.TestingT) *keeperFixture {
	f := &keeperFixture{}

	// gomock initializations
	ctrl := gomock.NewController(&testing.T{})
	supplyKeeper := crisistestutil.NewMockSupplyKeeper(ctrl)

	key := sdk.NewKVStoreKey(types.StoreKey)
	testCtx := testutil.DefaultContextWithDB(&testing.T{}, key, sdk.NewTransientStoreKey("transient_test"))
	encCfg := moduletestutil.MakeTestEncodingConfig(crisis.AppModuleBasic{})
	keeper := keeper.NewKeeper(encCfg.Codec, key, 5, supplyKeeper, "", "")

	f.ctx = testCtx.Ctx
	f.keeper = keeper
	f.authKeeper = supplyKeeper

	return f
}

func TestMsgVerifyInvariant(t *testing.T) {
	t.Parallel()
	f := initKeeperFixture(t)

	// default params
	constantFee := sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000))
	err := f.keeper.SetConstantFee(f.ctx, constantFee)
	assert.NilError(t, err)

	sender := sdk.AccAddress([]byte("addr1_______________"))

	f.authKeeper.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(2)
	f.keeper.RegisterRoute("bank", "total-supply", func(sdk.Context) (string, bool) { return "", false })

	testCases := []struct {
		name      string
		input     *types.MsgVerifyInvariant
		expErr    bool
		expErrMsg string
	}{
		{
			name: "empty sender not allowed",
			input: &types.MsgVerifyInvariant{
				Sender:              "",
				InvariantModuleName: "bank",
				InvariantRoute:      "total-supply",
			},
			expErr:    true,
			expErrMsg: "empty address string is not allowed",
		},
		{
			name: "invalid sender address",
			input: &types.MsgVerifyInvariant{
				Sender:              "invalid address",
				InvariantModuleName: "bank",
				InvariantRoute:      "total-supply",
			},
			expErr:    true,
			expErrMsg: "decoding bech32 failed",
		},
		{
			name: "unregistered invariant route",
			input: &types.MsgVerifyInvariant{
				Sender:              sender.String(),
				InvariantModuleName: "module",
				InvariantRoute:      "invalidroute",
			},
			expErr:    true,
			expErrMsg: "unknown invariant",
		},
		{
			name: "valid invariant",
			input: &types.MsgVerifyInvariant{
				Sender:              sender.String(),
				InvariantModuleName: "bank",
				InvariantRoute:      "total-supply",
			},
			expErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			_, err = f.keeper.VerifyInvariant(f.ctx, tc.input)
			if tc.expErr {
				assert.ErrorContains(t, err, tc.expErrMsg)
			} else {
				assert.NilError(t, err)
			}
		})
	}
}

func TestMsgUpdateParams(t *testing.T) {
	t.Parallel()
	f := initKeeperFixture(t)

	// default params
	constantFee := sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000))

	testCases := []struct {
		name      string
		input     *types.MsgUpdateParams
		expErr    bool
		expErrMsg string
	}{
		{
			name: "invalid authority",
			input: &types.MsgUpdateParams{
				Authority:   "invalid",
				ConstantFee: constantFee,
			},
			expErr:    true,
			expErrMsg: "invalid authority",
		},
		{
			name: "invalid constant fee",
			input: &types.MsgUpdateParams{
				Authority:   f.keeper.GetAuthority(),
				ConstantFee: sdk.Coin{},
			},
			expErr: true,
		},
		{
			name: "negative constant fee",
			input: &types.MsgUpdateParams{
				Authority:   f.keeper.GetAuthority(),
				ConstantFee: sdk.Coin{Denom: sdk.DefaultBondDenom, Amount: sdk.NewInt(-1000)},
			},
			expErr: true,
		},
		{
			name: "all good",
			input: &types.MsgUpdateParams{
				Authority:   f.keeper.GetAuthority(),
				ConstantFee: constantFee,
			},
			expErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			_, err := f.keeper.UpdateParams(f.ctx, tc.input)

			if tc.expErr {
				assert.ErrorContains(t, err, tc.expErrMsg)
			} else {
				assert.NilError(t, err)
			}
		})
	}
}
