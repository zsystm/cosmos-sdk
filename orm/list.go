package orm

import (
	"google.golang.org/protobuf/proto"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
)

type ListOptions struct {
	Reverse  bool
	UseIndex string
	Cursor   Cursor
	Start    proto.Message
	End      proto.Message
}

type Iterator interface {
	Next(proto.Message) (bool, error)
	Cursor() Cursor
	Close()

	mustEmbedUnimplementedIterator()
}

type Cursor []byte

type UnimplementedIterator struct{}

func (u UnimplementedIterator) mustEmbedUnimplementedIterator() {}

func (u UnimplementedIterator) Next(proto.Message) (bool, error) {
	return false, ormerrors.UnsupportedOperation
}

func (u UnimplementedIterator) Cursor() Cursor { return nil }

func (u UnimplementedIterator) Close() {}

var _ Iterator = UnimplementedIterator{}

type ErrIterator struct {
	UnimplementedIterator
	Err error
}

func (e ErrIterator) Cursor() Cursor { return nil }

func (e ErrIterator) Next(proto.Message) (bool, error) { return false, e.Err }

func (e ErrIterator) Close() {}

var _ Iterator = ErrIterator{}

type D struct {
}
