package groupclient

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	groupv1 "cosmossdk.io/api/cosmos/group/v1"
	"cosmossdk.io/client/v2"
)

func NewProposalClient(ctx context.Context, conn grpc.ClientConnInterface) (*ProposalClient, error) {
	msgClient := groupv1.NewMsgClient(conn)

	msgSubmitProposal := &groupv1.MsgSubmitProposal{}

	res, err := msgClient.SubmitProposal(ctx, msgSubmitProposal)
	pending, ok := err.(*clientv2.PendingTxResponse)
	if !ok {
		return nil, fmt.Errorf("expected %T, got %T", &clientv2.PendingTxResponse{}, err)
	}

	client := &ProposalClient{
		MsgSubmitProposal: msgSubmitProposal,
		res:               res,
	}

	pending.OnConfirmed(func() {
		// TODO
		//for _, result := range client.res.Results {
		//
		//}
	})

	return client, nil
}

type ProposalClient struct {
	*groupv1.MsgSubmitProposal
	res            *groupv1.MsgSubmitProposalResponse
	pendingReplies []interface{}
}

func (e *ProposalClient) Invoke(_ context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
	msg, ok := args.(proto.Message)
	if !ok {
		return fmt.Errorf("expected a proto.Message, got %T", args)
	}

	a, err := anypb.New(msg)
	if err != nil {
		return err
	}

	e.Messages = append(e.Messages, a)

	e.pendingReplies = append(e.pendingReplies, reply)

	return nil
}

func (e *ProposalClient) NewStream(context.Context, *grpc.StreamDesc, string, ...grpc.CallOption) (grpc.ClientStream, error) {
	return nil, fmt.Errorf("streaming not supported")
}
