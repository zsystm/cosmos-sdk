package clientv2

import (
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoregistry"

	"cosmossdk.io/tx/signing"
	"github.com/cosmos/cosmos-sdk/client/v2/keyring"
)

type Connection struct {
	nodeConn        grpc.ClientConnInterface
	signModeHandler signing.SignModeHandler
}

type ConnectionOptions struct {
	Keyring         keyring.Keyring
	SignModeHandler signing.SignModeHandler
}

func Connect(nodeUrl string, opts ConnectionOptions) (*Connection, error) {
	nodeConn, err := grpc.Dial(nodeUrl)
	if err != nil {
		return nil, err
	}
	return &Connection{nodeConn: nodeConn, signModeHandler: opts.SignModeHandler}, err
}

func (c *Connection) NewClient() *Client {
	return &Client{
		conn:          c,
		tx:            nil,
		txRaw:         nil,
		auxSignerData: nil,
	}
}

func (c *Connection) protoFiles() *protoregistry.Files {
	return protoregistry.GlobalFiles
}
