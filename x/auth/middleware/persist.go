package middleware

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

type persistHandler struct {
	next tx.Handler
}

// PersistTxMiddleware branch out MultiStore in DeliverTx and commit if no error returned.
func PersistTxMiddleware(txh tx.Handler) tx.Handler {
	return persistHandler{next: txh}
}

// CheckTx implements tx.Handler.CheckTx method.
func (sh persistHandler) CheckTx(ctx context.Context, req tx.Request, checkReq tx.RequestCheckTx) (tx.Response, tx.ResponseCheckTx, error) {
	// Do nothing during CheckTx.
	return sh.next.CheckTx(ctx, req, checkReq)
}

// DeliverTx implements tx.Handler.DeliverTx method.
func (sh persistHandler) DeliverTx(ctx context.Context, req tx.Request) (tx.Response, error) {
	// Create a new Context based off of the existing Context with a MultiStore branch
	// in case message processing fails. At this point, the MultiStore
	// is a branch of a branch.
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	runMsgCtx, msCache := cacheTxContext(sdkCtx, req.TxBytes)

	rsp, err := sh.next.DeliverTx(sdk.WrapSDKContext(runMsgCtx), req)
	if err != nil {
		msCache.Write()
	}

	return rsp, err
}

// SimulateTx implements tx.Handler.SimulateTx method.
func (sh persistHandler) SimulateTx(ctx context.Context, req tx.Request) (tx.Response, error) {
	// Create a new Context based off of the existing Context with a MultiStore branch
	// in case message processing fails. At this point, the MultiStore
	// is a branch of a branch.
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	runMsgCtx, msCache := cacheTxContext(sdkCtx, req.TxBytes)

	rsp, err := sh.next.SimulateTx(sdk.WrapSDKContext(runMsgCtx), req)
	if err != nil {
		msCache.Write()
	}

	return rsp, err
}

// cacheTxContext returns a new context based off of the provided context with
// a branched multi-store.
func cacheTxContext(sdkCtx sdk.Context, txBytes []byte) (sdk.Context, sdk.CacheMultiStore) {
	ms := sdkCtx.MultiStore()
	// TODO: https://github.com/cosmos/cosmos-sdk/issues/2824
	msCache := ms.CacheMultiStore()
	if msCache.TracingEnabled() {
		msCache = msCache.SetTracingContext(
			sdk.TraceContext(
				map[string]interface{}{
					"txHash": fmt.Sprintf("%X", tmhash.Sum(txBytes)),
				},
			),
		).(sdk.CacheMultiStore)
	}

	return sdkCtx.WithMultiStore(msCache), msCache
}
