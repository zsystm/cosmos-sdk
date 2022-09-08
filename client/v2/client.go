package clientv2

import (
	"context"

	signingv1beta1 "cosmossdk.io/api/cosmos/tx/signing/v1beta1"
	txv1beta1 "cosmossdk.io/api/cosmos/tx/v1beta1"
	"google.golang.org/grpc"
)

type Client struct {
	conn          *Connection
	txRaw         *txv1beta1.TxRaw
	auxSignerData *txv1beta1.AuxSignerData

	Tx *txv1beta1.Tx
}

var _ grpc.ClientConnInterface = &Client{}

func (c *Client) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) StartTx() error {
	// TODO check for an existing tx first
	c.Tx = &txv1beta1.Tx{}
	c.txRaw = &txv1beta1.TxRaw{}
	return nil
}

func (c *Client) DiscardTx() {
	c.Tx = nil
	c.txRaw = nil
}

// SignTx signs the transaction with the default sign mode.
func (c *Client) SignTx() error {
	return nil
}

// SignTxWithMode signs the transaction with the provided sign mode.
func (c *Client) SignTxWithMode(signingv1beta1.SignMode) error {
	return nil
}

func (c *Client) TxRaw() *txv1beta1.TxRaw {
	return c.txRaw
}

func (c *Client) AuxSignerData() *txv1beta1.AuxSignerData {
	return c.auxSignerData
}

func (c *Client) BroadcastTx() error {
	return nil
}

func (c *Client) WaitForConfirmation() error {
	return nil
}
