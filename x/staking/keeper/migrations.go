package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking/exported"
	v043 "github.com/cosmos/cosmos-sdk/x/staking/migrations/v043"
	v046 "github.com/cosmos/cosmos-sdk/x/staking/migrations/v046"
	v2 "github.com/cosmos/cosmos-sdk/x/staking/migrations/v2"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper         *Keeper
	legacySubspace exported.Subspace
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper *Keeper, ss exported.Subspace) Migrator {
	return Migrator{
		keeper:         keeper,
		legacySubspace: ss,
	}
}

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v043.MigrateStore(ctx, m.keeper.storeKey)
}

// Migrate2to3 migrates x/staking state from consensus version 2 to 3.
func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	return v046.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc, m.legacySubspace)
}

// Migrate3to4 migrates the x/staking module state from the consensus
// version 3 to version 4. Specifically, it takes the parameters that are currently stored
// and managed by the x/params modules and stores them directly into the x/staking
// module state.
func (m Migrator) Migrate3to4(ctx sdk.Context) error {
	return v2.Migrate(ctx, ctx.KVStore(m.keeper.storeKey), m.legacySubspace, m.keeper.cdc)
}
