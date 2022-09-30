package proc

import (
	"context"

	authnv1 "cosmossdk.io/api/cosmos/authn/v1"
)

// IncrementSequenceDecorator handles incrementing sequences of all signers.
// Use the IncrementSequenceDecorator decorator to prevent replay attacks. Note,
// there is need to execute IncrementSequenceDecorator on RecheckTx since
// BaseApp.Commit() will set the check state based on the latest header.
//
// NOTE: Since CheckTx and DeliverTx state are managed separately, subsequent and
// sequential txs orginating from the same account cannot be handled correctly in
// a reliable way unless sequence numbers are managed and tracked manually by a
// client. It is recommended to instead use multiple messages in a tx.
type IncrementSequenceDecorator struct {
	ak         authnv1.InternalClient
	authnState authnv1.StateStore
}

func NewIncrementSequenceDecorator(ak authnv1.InternalClient) IncrementSequenceDecorator {
	return IncrementSequenceDecorator{
		ak: ak,
	}
}

func (isd IncrementSequenceDecorator) ProcessTx(ctx context.Context, tx *TxInfo, simulate bool) (context.Context, error) {
	// increment sequence of all signers
	signers := GetSigners(tx)
	for _, signer := range signers {
		acc, err := isd.authnState.AccountTable().GetByAddress(ctx, signer)
		if err != nil {
			return nil, err
		}

		_, err = isd.ak.IncrementSeq(ctx, &authnv1.IncrementSeqRequest{AccountId: acc.Id})
		if err != nil {
			return nil, err
		}
	}

	return ctx, nil
}

func GetSigners(info *TxInfo) [][]byte {
	panic("TODO")
}
