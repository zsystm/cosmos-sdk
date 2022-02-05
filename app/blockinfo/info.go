package blockinfo

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service interface {
	GetBlockInfo(ctx context.Context) BlockInfo
}

type BlockInfo interface {
	Height() int64
	Time() *timestamppb.Timestamp
}
