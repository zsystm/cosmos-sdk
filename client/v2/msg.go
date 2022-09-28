package clientv2

import (
	"context"
	"fmt"
	"regexp"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"

	msgv1 "cosmossdk.io/api/cosmos/msg/v1"
)

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
		err = c.addMsg(msg)
		if err != nil {
			return err
		}

		c.pendingReplies = append(c.pendingReplies, reply)

		return &PendingTxResponse{}
	} else {
		return c.conn.nodeConn.Invoke(ctx, method, args, reply, opts...)
	}
}

func (c *Client) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	return c.conn.nodeConn.NewStream(ctx, desc, method, opts...)
}

func (c *Client) addMsg(msg proto.Message) error {
	msgDesc := msg.ProtoReflect().Descriptor()

	signers := proto.GetExtension(msgDesc.Options(), msgv1.E_Signer).([]string)
	if signers == nil {
		panic("TODO")
	}

	c.requiredSigners = append(c.requiredSigners, signers...)

	a, err := anypb.New(msg)
	if err != nil {
		return err
	}

	c.tx.Body.Messages = append(c.tx.Body.Messages, a)

	return nil
}
