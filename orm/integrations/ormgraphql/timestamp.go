package ormgraphql

import (
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

var timestampScalar = graphql.NewScalar(graphql.ScalarConfig{
	Name: "Timestamp",
	Serialize: func(value interface{}) interface{} {
		return strconv.FormatUint(uint64(value.(uint32)), 10)
	},
	ParseValue: func(value interface{}) interface{} {
		x, err := strconv.ParseUint(value.(string), 10, 32)
		if err != nil {
			return uint32(0)
		}
		return uint32(x)
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			x, err := strconv.ParseUint(valueAST.Value, 10, 32)
			if err != nil {
				return uint32(0)
			}
			return uint32(x)
		default:
			return nil
		}
	},
})
