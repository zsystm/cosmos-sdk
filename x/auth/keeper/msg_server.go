package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

type msgServer struct {
	AccountKeeper
}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper AccountKeeper) types.MsgServer {
	return &msgServer{AccountKeeper: keeper}
}

var _ types.MsgServer = msgServer{}

// UpdateParams is called by a predefined address.
// All fields must be set with the previous values, with the field(s) you would like to change, changed.
func (ms msgServer) UpdateParams(goCtx context.Context, msg *types.Params) (*types.ParamsUpdateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	ctx.

		// verification logic...

		ms.SetParams(ctx, *msg)

	return &types.ParamsUpdateResponse{}, nil
}
