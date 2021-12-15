package ormschema

import (
	"google.golang.org/protobuf/proto"

	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
)

type Schema interface {
	GetTable(message proto.Message) (ormtable.Table, error)
}
