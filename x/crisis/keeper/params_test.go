package keeper_test

import (
	"testing"

	"gotest.tools/v3/assert"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestParams(t *testing.T) {
	t.Parallel()
	f := initKeeperFixture(t)

	// default params
	constantFee := sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000))

	testCases := []struct {
		name        string
		constantFee sdk.Coin
		expErr      bool
		expErrMsg   string
	}{
		{
			name:        "invalid constant fee",
			constantFee: sdk.Coin{},
			expErr:      true,
		},
		{
			name:        "negative constant fee",
			constantFee: sdk.Coin{Denom: sdk.DefaultBondDenom, Amount: sdk.NewInt(-1000)},
			expErr:      true,
		},
		{
			name:        "all good",
			constantFee: constantFee,
			expErr:      false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			expected := f.keeper.GetConstantFee(f.ctx)
			err := f.keeper.SetConstantFee(f.ctx, tc.constantFee)

			if tc.expErr {
				assert.ErrorContains(t, err, tc.expErrMsg)
			} else {
				expected = tc.constantFee
				assert.NilError(t, err)
			}

			params := f.keeper.GetConstantFee(f.ctx)

			assert.DeepEqual(t, expected, params)
		})
	}
}
