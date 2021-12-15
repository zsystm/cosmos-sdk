package ormgraphql

import (
	"encoding/base64"

	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

var cursor = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "Cursor",
	Description: "",
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case ormtable.Cursor:
			return base64.RawURLEncoding.EncodeToString(value)
		default:
			return nil
		}
	},
	ParseValue: func(value interface{}) interface{} {
		switch value := value.(type) {
		case string:
			return ormtable.Cursor(decodeBase64Scalar(value))
		case *string:
			return ormtable.Cursor(decodeBase64Scalar(*value))
		default:
			return nil
		}
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			return ormtable.Cursor(decodeBase64Scalar(valueAST.Value))
		default:
			return nil
		}
	},
})
