package keeper

import (
	"context"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
)

type BaseContext interface {
	context.Context

	BlockHeader() tmproto.Header
	TxBytes() []byte
	TxHash() []byte
	EventManager() sdk.EventManager
}

type PrepareContext interface {
	BaseContext

	GetRef(key []byte) KVRef
	GetIteratorRef(key []byte) IteratorRef
	Exec(func(ExecContext) error)
}

type KVRef interface {
	Value(ExecContext) []byte
	SetValue(ExecContext, []byte)
}

type IteratorRef interface {
	Next(ExecContext) bool
	Key() []byte
	Value() []byte
	SetValue([]byte)
}

type ExecContext interface {
	BaseContext
}

type parallelMsgServerImpl struct {
}

func AccountBalanceKey(addr sdk.AccAddress, denom string) []byte {
	panic("TODO")
}

var _ MsgServerParallel = parallelMsgServerImpl{}

type MsgServerParallel interface {
	// Send defines a method for sending coins from one account to another account.
	Send(ctx PrepareContext, request *types.MsgSend, setResponse func(*types.MsgSendResponse)) error
	// MultiSend defines a method for sending coins from some accounts to other accounts.
	MultiSend(ctx PrepareContext, request *types.MsgMultiSend, setResponse func(response *types.MsgMultiSendResponse)) error
}

func (m parallelMsgServerImpl) Send(ctx PrepareContext, request *types.MsgSend, setResponse func(*types.MsgSendResponse)) error {
	// NOTE: can read last block's state synchronously
	//if err := k.SendEnabledCoins(ctx, msg.Amount...); err != nil {
	//	return nil, err
	//}

	from, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return err
	}
	to, err := sdk.AccAddressFromBech32(msg.ToAddress)
	if err != nil {
		return err
	}

	// NOTE: can read last block's state synchronously
	//if k.BlockedAddr(to) {
	//	return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to receive funds", msg.ToAddress)
	//}

	for _, coin := range msg.Amount {
		fromBalance := ctx.GetRef(AccountBalanceKey(from, coin.Denom))
		toBalance := ctx.GetRef(AccountBalanceKey(to, coin.Denom))
		ctx.Exec(func(ctx ExecContext) error {
		})
	}

	return nil
}

func (m parallelMsgServerImpl) MultiSend(ctx PrepareContext, request *types.MsgMultiSend, setResponse func(response *types.MsgMultiSendResponse)) error {
	panic("implement me")
}

type MsgClientParallel interface {
	// Send defines a method for sending coins from one account to another account.
	Send(ctx PrepareContext, request *types.MsgSend) (MsgSendExecutor, error)
	// MultiSend defines a method for sending coins from some accounts to other accounts.
	MultiSend(ctx PrepareContext, request *types.MsgMultiSend) (MsgMultiSendExecutor, error)
}

type MsgSendExecutor func(ExecContext) (*types.MsgSendResponse, error)

type MsgMultiSendExecutor func(ExecContext) (*types.MsgSendResponse, error)
