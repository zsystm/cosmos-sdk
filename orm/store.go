package orm

import (
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"

	"google.golang.org/protobuf/proto"
)

type ReadStore interface {
	Has(...proto.Message) bool
	Get(...proto.Message) (found bool, err error)
	List(condition proto.Message, options *ListOptions) Iterator

	mustEmbedUnimplementedReadStore()
}

type Store interface {
	ReadStore

	Create(...proto.Message) error
	Save(...proto.Message) error
	Delete(...proto.Message) error

	mustEmbedUnimplementedStore()
}

type UnimplementedReadStore struct{}

func (UnimplementedReadStore) mustEmbedUnimplementedReadStore() {}

func (u UnimplementedReadStore) Has(...proto.Message) bool {
	return false
}

func (u UnimplementedReadStore) Get(...proto.Message) (found bool, err error) {
	return false, ormerrors.UnsupportedOperation
}

func (u UnimplementedReadStore) List(proto.Message, *ListOptions) Iterator {
	return ErrIterator{Err: ormerrors.UnsupportedOperation}
}

var _ ReadStore = UnimplementedReadStore{}

type UnimplementedStore struct {
	UnimplementedReadStore
}

func (u UnimplementedStore) mustEmbedUnimplementedStore() {}

func (u UnimplementedStore) Create(...proto.Message) error {
	return ormerrors.UnsupportedOperation
}

func (u UnimplementedStore) Save(...proto.Message) error {
	return ormerrors.UnsupportedOperation
}

func (u UnimplementedStore) Delete(...proto.Message) error {
	return ormerrors.UnsupportedOperation
}

var _ Store = UnimplementedStore{}
