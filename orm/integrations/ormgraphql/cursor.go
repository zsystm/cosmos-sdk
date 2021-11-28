package ormgraphql

import (
	"encoding/base64"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"

	"github.com/cosmos/cosmos-sdk/orm/model/ormiterator"
)

var cursor = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "Cursor",
	Description: "",
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case ormiterator.Cursor:
			return base64.RawURLEncoding.EncodeToString(value)
		default:
			return nil
		}
	},
	ParseValue: func(value interface{}) interface{} {
		switch value := value.(type) {
		case string:
			return ormiterator.Cursor(decodeBase64Scalar(value))
		case *string:
			return ormiterator.Cursor(decodeBase64Scalar(*value))
		default:
			return nil
		}
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			return ormiterator.Cursor(decodeBase64Scalar(valueAST.Value))
		default:
			return nil
		}
	},
})

func decodeBase64Scalar(value string) []byte {
	bytes, err := base64.RawURLEncoding.DecodeString(value)
	if err != nil {
		return nil
	}
	return bytes
}
