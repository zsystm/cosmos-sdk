package graphql

import (
	"github.com/graphql-go/graphql"

	"github.com/cosmos/cosmos-sdk/orm/client/graphql/internal/value"
)

var pageInfo = graphql.NewObject(graphql.ObjectConfig{
	Name:       "PageInfo",
	Interfaces: nil,
	Fields: graphql.Fields{
		"hasNextPage": {
			Name:        "hasNextPage",
			Type:        graphql.Boolean,
			Resolve:     nil,
			Subscribe:   nil,
			Description: "",
		},
		"hasPreviousPage": {
			Name:        "hasPreviousPage",
			Type:        graphql.Boolean,
			Resolve:     nil,
			Subscribe:   nil,
			Description: "",
		},
		"startCursor": {
			Name:        "startCursor",
			Type:        value.Cursor,
			Resolve:     nil,
			Subscribe:   nil,
			Description: "",
		},
		"endCursor": {
			Name:        "endCursor",
			Type:        value.Cursor,
			Resolve:     nil,
			Subscribe:   nil,
			Description: "",
		},
	},
	IsTypeOf:    nil,
	Description: "Information about pagination in a connection.",
})
