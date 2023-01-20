package testutil

import (
	"fmt"
	"testing"

	"cosmossdk.io/depinject"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil/sims"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/cosmos/cosmos-sdk/x/staking/keeper"
)

func TestSimApp(depModules []string, stakingkeeper *keeper.Keeper, accountKeeper authkeeper.AccountKeeper, bankkeeper bankkeeper.Keeper, config depinject.Config) *runtime.App {
	// if c, ok := modules.(keeper.Keeper); ok { // type assert on it
	// 	fmt.Println(c.Name)
	// }
	fmt.Println("=====")
	// fmt.Println(replaceModules)
	fmt.Println("-----------")
	app, err := sims.SetupWithConfiguration(config, sims.DefaultStartUpConfig(), &stakingkeeper, &accountKeeper, &bankkeeper)
	assert.NilError(&testing.T{}, err)

	return app
}
