package testing

import (
	"context"

	"google.golang.org/grpc"

	"cosmossdk.io/core/genesisjson"
)

type Fixture interface {
	Context() context.Context
	QueryConn() grpc.ClientConnInterface
	MsgConn() grpc.ClientConnInterface

	SignerAddress(i uint) string

	BeginBlock() error
	EndBlock() error

	InitGenesis(genesisjson.Source) error
	ExportGenesis(genesisjson.Target) error
}
