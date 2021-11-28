package ormgraphql

import (
	"fmt"

	"github.com/graphql-go/graphql"

	"github.com/cosmos/cosmos-sdk/orm/model/ormindex"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
)

func (b Builder) buildUniqueIndex(table ormtable.Table, index ormindex.UniqueIndex, fields graphql.Fields) error {
	fieldName := fmt.Sprintf(
		"get%sby%s",
		table.MessageType().Descriptor().Name(),
		//panic("TODO")
	)
	fields[fieldName] = &graphql.Field{
		Name:              fieldName,
		Type:              nil,
		Args:              nil,
		Resolve:           nil,
		Subscribe:         nil,
		DeprecationReason: "",
		Description:       "",
	}
	return nil
}
