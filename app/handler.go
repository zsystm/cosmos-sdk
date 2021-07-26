package app

import (
	"context"
	"encoding/json"

	abci "github.com/tendermint/tendermint/abci/types"
	"google.golang.org/grpc"

	"github.com/cosmos/cosmos-sdk/codec"
)

type Handler struct {
	InitGenesis   func(context.Context, codec.JSONCodec, json.RawMessage) []abci.ValidatorUpdate
	BeginBlocker  func(context.Context, abci.RequestBeginBlock)
	EndBlocker    func(context.Context, abci.RequestEndBlock) []abci.ValidatorUpdate
	MsgServices   []ServiceImpl
	QueryServices []ServiceImpl
}

type ServiceImpl struct {
	Desc *grpc.ServiceDesc
	Impl interface{}
}
