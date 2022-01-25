package op

import (
	"context"

	"google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
)

type Filter struct {
	Func func(proto.Message) bool
	Op   Op
}

func (f Filter) Exec(ctx context.Context) (ormtable.Iterator, error) {
	it, err := f.Op.Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &filterIterator{
		it:     it,
		filter: f.Func,
	}, nil
}

type filterIterator struct {
	ormtable.UnimplementedIterator
	it     ormtable.Iterator
	err    error
	cur    proto.Message
	filter func(proto.Message) bool
}

func (f filterIterator) Next() bool {
	if f.err != nil {
		return false
	}

	for f.it.Next() {
		f.cur, f.err = f.it.GetMessage()
		if f.err != nil {
			return false
		}
		if f.filter(f.cur) {
			return true
		}
	}
	return false
}

func (f filterIterator) Keys() (indexKey, primaryKey []protoreflect.Value, err error) {
	return f.it.Keys()
}

func (f filterIterator) UnmarshalMessage(message proto.Message) error {
	return f.it.UnmarshalMessage(message)
}

func (f filterIterator) GetMessage() (proto.Message, error) {
	return f.cur, f.err
}

func (f filterIterator) Cursor() ormlist.CursorT {
	return f.it.Cursor()
}

func (f filterIterator) Close() {
	f.it.Close()
}

var _ ormtable.Iterator = filterIterator{}
