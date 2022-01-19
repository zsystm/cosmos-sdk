package server

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	bankv1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/bank/v1beta1"
	bankv2alpha1 "github.com/cosmos/cosmos-sdk/api/cosmos/bank/v2alpha1"
)

var _ bankv1beta1.MsgServer = &server{}

func (s server) Send(ctx context.Context, send *bankv1beta1.MsgSend) (*bankv1beta1.MsgSendResponse, error) {
	fromAddressBz, err := s.addressCodec.StringToBytes(send.FromAddress)
	if err != nil {
		return nil, err
	}

	toAddressBz, err := s.addressCodec.StringToBytes(send.ToAddress)
	if err != nil {
		return nil, err
	}

	for _, coin := range send.Amount {
		amount, ok := sdk.NewIntFromString(coin.Amount)
		if !ok {
			return nil, fmt.Errorf("invalid amount %s", coin.Amount)
		}

		// get from balance
		var fromBalance, toBalance bankv2alpha1.Balance
		found, err := s.balanceAddressDenomIndex.Get(ctx, &fromBalance, fromAddressBz, coin.Denom)
		if err != nil {
			return nil, err
		}

		if !found {
			return nil, fmt.Errorf("no from balance for %s", coin.Denom)
		}

		fromAmount, ok := sdk.NewIntFromString(fromBalance.Amount)
		if !ok {
			return nil, fmt.Errorf("invalid amount %s", fromBalance.Amount)
		}

		// update from balance
		newFromAmount := fromAmount.Sub(amount)
		if newFromAmount.IsNegative() {
			return nil, errors.ErrInsufficientFunds
		}

		fromBalance.Amount = newFromAmount.String()
		err = s.balanceTable.Save(ctx, &fromBalance)
		if err != nil {
			return nil, err
		}

		// get to balance
		found, err = s.balanceAddressDenomIndex.Get(ctx, &toBalance, toAddressBz, coin.Denom)
		if err != nil {
			return nil, err
		}

		if !found {
			toBalance.Address = toAddressBz
			toBalance.Denom = coin.Denom
			toBalance.Amount = "0"
		}

		toAmount, ok := sdk.NewIntFromString(toBalance.Amount)
		if !ok {
			return nil, fmt.Errorf("invalid amount %s", toBalance.Amount)
		}

		// update to balance
		newToAmount := toAmount.Add(amount)
		toBalance.Amount = newToAmount.String()
		err = s.balanceTable.Save(ctx, &toBalance)
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
