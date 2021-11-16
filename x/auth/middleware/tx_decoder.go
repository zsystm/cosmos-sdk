package middleware

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
)

func NewTxDecoderMiddleware(txDecoder sdk.TxDecoder) {
	return
}

type txDecoderHandler struct {
	txDecoder sdk.TxDecoder
	next      tx.Handler
}

func (txh *txDecoderHandler) DeliverTx() {
	sdkTx, err := txh.txDecoder(req.TxBytes)

	txh.next(tx.Request{TxBytes: req.TxBytes, Tx: sdkTx})
}
