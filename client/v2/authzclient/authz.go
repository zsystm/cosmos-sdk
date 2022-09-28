package authzclient

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	authzv1beta1 "cosmossdk.io/api/cosmos/authz/v1beta1"
	clientv2 "cosmossdk.io/client/v2"
)

func NewExecClient(ctx context.Context, conn grpc.ClientConnInterface, grantee string) (*ExecClient, error) {
	msgClient := authzv1beta1.NewMsgClient(conn)

	msgExec := &authzv1beta1.MsgExec{
		Grantee: grantee,
	}

	res, err := msgClient.Exec(ctx, msgExec)
	pending, ok := err.(*clientv2.PendingTxResponse)
	if !ok {
		return nil, fmt.Errorf("expected %T, got %T", &clientv2.PendingTxResponse{}, err)
	}

	client := &ExecClient{
		MsgExec: msgExec,
		res:     res,
	}

	pending.OnConfirmed(func() {
		// TODO
		//for _, result := range client.res.Results {
		//
		//}
	})

	return client, nil
}

type ExecClient struct {
	*authzv1beta1.MsgExec
	res            *authzv1beta1.MsgExecResponse
	pendingReplies []interface{}
}

func (e *ExecClient) Invoke(_ context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
	msg, ok := args.(proto.Message)
	if !ok {
		return fmt.Errorf("expected a proto.Message, got %T", args)
	}

	a, err := anypb.New(msg)
	if err != nil {
		return err
	}

	e.Msgs = append(e.Msgs, a)

	e.pendingReplies = append(e.pendingReplies, reply)

	return nil
}

func (e *ExecClient) NewStream(context.Context, *grpc.StreamDesc, string, ...grpc.CallOption) (grpc.ClientStream, error) {
	return nil, fmt.Errorf("streaming not supported")
}
