package clientv2

import (
	"google.golang.org/grpc"

	authzv1beta1 "cosmossdk.io/api/cosmos/authz/v1beta1"
	signingv1beta1 "cosmossdk.io/api/cosmos/tx/signing/v1beta1"
	txv1beta1 "cosmossdk.io/api/cosmos/tx/v1beta1"
)

type Client struct {
	conn            *Connection
	txRaw           *txv1beta1.TxRaw
	auxSignerData   *txv1beta1.AuxSignerData
	requiredSigners []string

	tx             *txv1beta1.Tx
	pendingReplies []interface{}
	authzExec      *authzv1beta1.MsgExec
}

var _ grpc.ClientConnInterface = &Client{}

func (c *Client) DiscardTx() {
	// TODO start new tx
}

// SignTx signs the transaction with the default sign mode.
func (c *Client) SignTx() error {
	//return c.SignTxWithMode(c.conn.signModeHandler.DefaultMode())
	panic("TODO")
}

// SignTxWithMode signs the transaction with the provided sign mode.
func (c *Client) SignTxWithMode(signingv1beta1.SignMode) error {
	// TODO check if signers have acct num & seq, if not retrieve
	panic("TODO")
}

func (c *Client) Tx() *txv1beta1.Tx {
	return c.tx
}

func (c *Client) TxRaw() *txv1beta1.TxRaw {
	return c.txRaw
}

func (c *Client) AuxSignerData() *txv1beta1.AuxSignerData {
	return c.auxSignerData
}

func (c *Client) BroadcastTx() (*PendingTxResponse, error) {
	return &PendingTxResponse{}, nil
}

type PendingTxResponse struct{}

func (p *PendingTxResponse) Error() string {
	return "pending tx response"
}

func (p *PendingTxResponse) Is(target error) bool {
	_, ok := target.(*PendingTxResponse)
	return ok
}

func (p *PendingTxResponse) OnConfirmed(f func()) {

}

func (p *PendingTxResponse) AwaitConfirmation() error {
	panic("TODO")
}

var _ error = &PendingTxResponse{}
