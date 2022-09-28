package clientv2

import (
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoregistry"

	"cosmossdk.io/client/v2/keyring"
	"cosmossdk.io/tx/signing"
)

type Connection struct {
	ConnectionOptions
	nodeConn grpc.ClientConnInterface
}

type ConnectionOptions struct {
	Keyring         keyring.Keyring
	SignModeHandler signing.SignModeHandler
}

func (opts ConnectionOptions) Connect(nodeUrl string) (*Connection, error) {
	nodeConn, err := grpc.Dial(nodeUrl)
	if err != nil {
		return nil, err
	}
	return &Connection{nodeConn: nodeConn, ConnectionOptions: opts}, err
}

func Connect(nodeUrl string) (*Connection, error) {
	return ConnectionOptions{}.Connect(nodeUrl)
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
