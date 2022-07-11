package keeper

import (
	"time"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
)

// UnbondingTime
func (k Keeper) UnbondingTime(ctx sdk.Context) (res time.Duration) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyUnbondingTime)
	if bz == nil {
		return res
	}
	var params types.Params
	k.cdc.MustUnmarshal(bz, &params)
	return params.UnbondingTime
}

// MaxValidators - Maximum number of validators
func (k Keeper) MaxValidators(ctx sdk.Context) (res uint32) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyMaxValidators)
	if bz == nil {
		return res
	}
	var params types.Params
	k.cdc.MustUnmarshal(bz, &params)
	return params.MaxValidators
}

// MaxEntries - Maximum number of simultaneous unbonding
// delegations or redelegations (per pair/trio)
func (k Keeper) MaxEntries(ctx sdk.Context) (res uint32) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyMaxEntries)
	if bz == nil {
		return res
	}
	var params types.Params
	k.cdc.MustUnmarshal(bz, &params)
	return params.MaxEntries
}

// HistoricalEntries = number of historical info entries
// to persist in store
func (k Keeper) HistoricalEntries(ctx sdk.Context) (res uint32) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyHistoricalEntries)
	if bz == nil {
		return res
	}
	var params types.Params
	k.cdc.MustUnmarshal(bz, &params)
	return params.HistoricalEntries
}

// BondDenom - Bondable coin denomination
func (k Keeper) BondDenom(ctx sdk.Context) (res string) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyBondDenom)
	if bz == nil {
		return res
	}
	var params types.Params
	k.cdc.MustUnmarshal(bz, &params)
	return params.BondDenom
}

// PowerReduction - is the amount of staking tokens required for 1 unit of consensus-engine power.
// Currently, this returns a global variable that the app developer can tweak.
// TODO: we might turn this into an on-chain param:
// https://github.com/cosmos/cosmos-sdk/issues/8365
func (k Keeper) PowerReduction(ctx sdk.Context) math.Int {
	return sdk.DefaultPowerReduction
}

// MinCommissionRate - Minimum validator commission rate
func (k Keeper) MinCommissionRate(ctx sdk.Context) (res sdk.Dec) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyMinCommissionRate)
	if bz == nil {
		return res
	}
	var params types.Params
	k.cdc.MustUnmarshal(bz, &params)
	return params.MinCommissionRate
}

// Get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	return types.NewParams(
		k.UnbondingTime(ctx),
		k.MaxValidators(ctx),
		k.MaxEntries(ctx),
		k.HistoricalEntries(ctx),
		k.BondDenom(ctx),
		k.MinCommissionRate(ctx),
	)
}

// set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) error {
	if err := params.Validate(); err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&params)
	store.Set(types.ParamsKey, bz)

	return nil
}
