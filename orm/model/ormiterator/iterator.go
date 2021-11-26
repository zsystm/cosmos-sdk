package ormiterator

import (
	"google.golang.org/protobuf/proto"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
)

type Iterator interface {
	mustEmbedUnimplementedIterator()

	Next(proto.Message) (bool, error)

	// TODO
	//Valid() bool
	//Next()
	//IndexKey() ([]protoreflect.Value, error)
	//PrimaryKey() ([]protoreflect.Value, error)
	//Value(proto.Message) error

	Cursor() Cursor
	Close()
}

type Cursor []byte

type UnimplementedIterator struct{}

func (u UnimplementedIterator) mustEmbedUnimplementedIterator() {}

func (u UnimplementedIterator) Next(proto.Message) (bool, error) {
	return false, ormerrors.UnsupportedOperation
}

func (u UnimplementedIterator) Cursor() Cursor { return nil }

func (u UnimplementedIterator) Close() {
	panic("implement me")
}

var _ Iterator = UnimplementedIterator{}

type ErrIterator struct {
	UnimplementedIterator
	Err error
}

func (e ErrIterator) Cursor() Cursor { return nil }

func (e ErrIterator) Next(proto.Message) (bool, error) { return false, e.Err }

func (e ErrIterator) Close() {}

var _ Iterator = ErrIterator{}
