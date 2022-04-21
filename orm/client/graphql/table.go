package graphql

import (
	"github.com/graphql-go/graphql"

	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
)

func (b Builder) buildTable(table ormtable.Table) graphql.Fields {
	fields := map[string]*graphql.Field{}

	for _, index := range table.Indexes() {
		if _, ok := index.(ormtable.UniqueIndex); ok {
			panic("TODO")
		}
		panic("TODO")
	}

	return fields
}
