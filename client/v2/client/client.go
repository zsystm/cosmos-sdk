package client

import (
	"context"

	txv1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/tx/v1beta1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

type Client struct {
	txBody   *txv1beta1.TxBody
	authInfo *txv1beta1.AuthInfo
}

func (c Client) Context() context.Context {
	panic("TODO")
}

func (c Client) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
	//TODO implement me
	panic("implement me")
}

func (c Client) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	//TODO implement me
	panic("implement me")
}

var _ grpc.ClientConnInterface = Client{}

func (c *Client) BeginTx() error {
	panic("TODO")
}

func (c *Client) AddMsg(message proto.Message) {}

func (c *Client) SignTx() error {
	panic("TODO")
}

func (c Client) BroadcastTx() error {
	panic("TODO")
}
