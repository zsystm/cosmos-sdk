package orm

import (
	"context"

	"google.golang.org/protobuf/proto"

	"github.com/cosmos/cosmos-sdk/orm/model/ormschema"
)

type DB interface {
	OpenRead(context.Context) (*ReadDBConnection, error)
	Open(context.Context) (*DBConnection, error)
}

type db struct {
	ormschema.DB
}

func (d db) GetTable(message proto.Message) (Table, error) {
	panic("TODO")
}

type Table interface {
	GetIndex(fieldNames string) (Index, error)
	GetUniqueIndex(fieldNames string) (UniqueIndex, error)
}

type Index interface {
}

type UniqueIndex interface {
	Index
	Get(conn Conn, message proto.Message, fieldValues ...interface{}) (found bool, err error)
	Has(conn Conn, message proto.Message, fieldValues ...interface{})
}

type Conn interface {
	todo()
}
