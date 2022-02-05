package app

import (
	"context"

	"google.golang.org/grpc"

	"github.com/cosmos/cosmos-sdk/container"
)

type Handler struct {
	MsgServices       []ServiceImpl
	QueryServices     []ServiceImpl
	BasicBeginBlocker func(context.Context) error
	BasicEndBlocker   func(context.Context) error
}

type ServiceImpl struct {
	Desc *grpc.ServiceDesc
	Impl interface{}
}

func (h Handler) IsOnePerModuleType() {}

var _ container.OnePerModuleType = Handler{}
