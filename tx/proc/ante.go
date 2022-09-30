package proc

import (
	"context"

	txv1beta1 "cosmossdk.io/api/cosmos/tx/v1beta1"
)

type TxInfo struct {
	TxBytes []byte
	TxRaw   *txv1beta1.TxRaw
	Tx      *txv1beta1.Tx
}

type TxProcessor interface {
	ProcessTx(ctx context.Context, tx *TxInfo) (newCtx context.Context, err error)
}
