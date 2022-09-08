package clientv2

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"

	msgv1 "cosmossdk.io/api/cosmos/msg/v1"
	signingv1beta1 "cosmossdk.io/api/cosmos/tx/signing/v1beta1"
	txv1beta1 "cosmossdk.io/api/cosmos/tx/v1beta1"
)

var Pending = errors.New("pending")

type Client struct {
	conn            *Connection
	txRaw           *txv1beta1.TxRaw
	auxSignerData   *txv1beta1.AuxSignerData
	requiredSigners []string

	tx             *txv1beta1.Tx
	pendingReplies []interface{}
}

var _ grpc.ClientConnInterface = &Client{}

var methodRegex = regexp.MustCompile(`/(.+)/(.+)`)

func (c *Client) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
	matches := methodRegex.FindStringSubmatch(method)
	if len(matches) < 3 {
		panic("TODO")
	}

	serviceName := matches[1]
	desc, err := c.conn.protoFiles().FindDescriptorByName(protoreflect.FullName(serviceName))
	if err != nil {
		return err
	}

	serviceDesc, ok := desc.(protoreflect.ServiceDescriptor)
	if !ok {
		return fmt.Errorf("expected service")
	}

	isMsgService := proto.GetExtension(serviceDesc.Options(), msgv1.E_Service).(bool)
	if isMsgService {
		if c.tx == nil {
			panic("TODO")
		}

		if len(opts) > 0 {
			return fmt.Errorf("no options are available for tx msg's")
		}

		msg := args.(proto.Message)
		msgDesc := msg.ProtoReflect().Descriptor()

		signers := proto.GetExtension(msgDesc.Options(), msgv1.E_Signer).([]string)
		if signers == nil {
			panic("TODO")
		}

		c.requiredSigners = append(c.requiredSigners, signers...)

		any, err := anypb.New(args.(proto.Message))
		if err != nil {
			return err
		}

		c.tx.Body.Messages = append(c.tx.Body.Messages, any)

		c.pendingReplies = append(c.pendingReplies, reply)

		return Pending
	} else {
		return c.conn.nodeConn.Invoke(ctx, method, args, reply, opts...)
	}
}

func (c *Client) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	return c.conn.nodeConn.NewStream(ctx, desc, method, opts...)
}

func (c *Client) DiscardTx() {
	// TODO start new tx
}

// SignTx signs the transaction with the default sign mode.
func (c *Client) SignTx() error {
	return nil
}

// SignTxWithMode signs the transaction with the provided sign mode.
func (c *Client) SignTxWithMode(signingv1beta1.SignMode) error {
	return nil
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

func (p *PendingTxResponse) WaitForConfirmation() error {
	return nil
}
