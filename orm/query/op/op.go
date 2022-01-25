package op

import (
	"context"

	"google.golang.org/protobuf/proto"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
)

type Op interface {
	Exec(context.Context) (ormtable.Iterator, error)
}

type IndexScan struct {
	Index   ormtable.Index
	Options []ormlist.Option
}

func (i IndexScan) Exec(ctx context.Context) (ormtable.Iterator, error) {
	return i.Index.Iterator(ctx, i.Options...)
}

type IndexIntersection struct {
	IndexScans []*IndexScan
}

func (i IndexIntersection) Exec(ctx context.Context) (ormtable.Iterator, error) {
	//TODO implement me
	panic("implement me")
}

type Sort struct {
	Cmp func(a, b proto.Message) int
	Op  Op
}

func (s Sort) Exec(ctx context.Context) (ormtable.Iterator, error) {
	//TODO implement me
	panic("implement me")
}

type OffsetLimit struct {
	Offset, Limit int
	Op            Op
}

func (o OffsetLimit) Exec(ctx context.Context) (ormtable.Iterator, error) {
	//TODO implement me
	panic("implement me")
}

var _, _, _, _, _ Op = IndexScan{}, IndexIntersection{}, Sort{}, Filter{}, OffsetLimit{}
