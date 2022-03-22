package testutil

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestIntegrationTestSuite(t *testing.T) {
	cfg := network.DefaultConfig()
	cfg.NumValidators = 2
	suite.Run(t, NewIntegrationTestSuite(cfg))
}

func TestUnbondTestSuite(t *testing.T) {
	cfg := network.DefaultConfig()
	cfg.NumValidators = 2

	genesisState := types.DefaultGenesisState()

	// change the unbonding period to 5 seconds.
	genesisState.Params.UnbondingTime = 5 * time.Second

	bz, err := cfg.Codec.MarshalJSON(genesisState)
	require.NoError(t, err)
	cfg.GenesisState["staking"] = bz
	suite.Run(t, NewTestSuiteUnbond(cfg))
}
