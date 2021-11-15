package keeper

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

var emptyParams = types.Params{}

// SetParams sets the auth module's parameters.
func (ak AccountKeeper) SetParams(ctx sdk.Context, params types.Params) {
	store := ctx.KVStore(ak.key)
	bz, err := ak.cdc.Marshal(&params)
	if err != nil {
		panic(err)
	}

	store.Set(types.ParamStoreKey, bz)
}

// GetParams gets the auth module's parameters.
func (ak AccountKeeper) GetParams(ctx sdk.Context) (types.Params, error) {
	store := ctx.KVStore(ak.key)
	bz := store.Get(types.ParamStoreKey)
	if bz == nil {
		return emptyParams, errors.New("unable to find params")
	}

	params := types.Params{}

	if err := ak.cdc.Unmarshal(bz, &params); err != nil {
		return emptyParams, err
	}

	return params, nil
}
