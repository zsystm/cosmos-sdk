package blockinfo

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service struct{}

type BlockInfo interface {
	Height() int64
	Time() *timestamppb.Timestamp

	private()
}

func (s *Service) GetBlockInfo(ctx context.Context) BlockInfo {
	panic("TODO")
}
