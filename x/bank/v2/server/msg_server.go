package server

import (
	"context"
	"fmt"

	bankv1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/bank/v1beta1"
	bankv2alpha1 "github.com/cosmos/cosmos-sdk/api/cosmos/bank/v2alpha1"
)

var _ bankv1beta1.MsgServer = &server{}

func (s server) Send(ctx context.Context, send *bankv1beta1.MsgSend) (*bankv1beta1.MsgSendResponse, error) {
	db, err := s.dbConnection.Open(ctx)
	if err != nil {
		return nil, err
	}

	addressBz, err := s.addressCodec.StringToBytes(send.FromAddress)
	if err != nil {
		return nil, err
	}

	for _, coin := range send.Amount {
		var fromBalance, toBalance bankv2alpha1.Balance
		found, err := db.Get(&fromBalance, addressDenomFields, addressBz, coin.Denom)
		if err != nil {
			return nil, err
		}

		if !found {
			return nil, fmt.Errorf("no balance for %s", coin.Denom)
		}

		// TODO check balance
		// TODO retrieve toBalance
		// TODO do math
		var newFromAmount, newToAmount string
		fromBalance.Amount = newFromAmount
		toBalance.Amount = newToAmount

		err = db.Save(&fromBalance)
		if err != nil {
			return nil, err
		}

		err = db.Save(&toBalance)
		if err != nil {
			return nil, err
		}
	}

	return &bankv1beta1.MsgSendResponse{}, nil
}

func (s server) MultiSend(ctx context.Context, send *bankv1beta1.MsgMultiSend) (*bankv1beta1.MsgMultiSendResponse, error) {
	//TODO implement me
	panic("implement me")
}
