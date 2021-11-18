package txmiddleware

import (
	"context"

	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
)

type circuitBreakerMiddleware struct {
	keeper crisiskeeper.Keeper
	next   tx.Handler
}

var (
	_ tx.Handler = circuitBreakerMiddleware{}
)

func (c circuitBreakerMiddleware) CheckTx(goCtx context.Context, tx sdk.Tx, req abci.RequestCheckTx) (abci.ResponseCheckTx, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if c.keeper.IsCircuitBreakerTripped(ctx) {
	}
	return c.next.CheckTx(goCtx, tx, req)
}

func (c circuitBreakerMiddleware) DeliverTx(ctx context.Context, tx sdk.Tx, req abci.RequestDeliverTx) (abci.ResponseDeliverTx, error) {
	return c.next.DeliverTx(ctx, tx, req)
}

func (c circuitBreakerMiddleware) SimulateTx(ctx context.Context, tx sdk.Tx, req tx.RequestSimulateTx) (tx.ResponseSimulateTx, error) {
	return c.next.SimulateTx(ctx, tx, req)
}
