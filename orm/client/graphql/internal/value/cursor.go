package value

import (
	"encoding/base64"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

var Cursor = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "Cursor",
	Description: "",
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case ormlist.CursorT:
			return base64.RawURLEncoding.EncodeToString(value)
		default:
			return nil
		}
	},
	ParseValue: func(value interface{}) interface{} {
		switch value := value.(type) {
		case string:
			return ormlist.CursorT(DecodeBase64Scalar(value))
		case *string:
			return ormlist.CursorT(DecodeBase64Scalar(*value))
		default:
			return nil
		}
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			return ormlist.CursorT(DecodeBase64Scalar(valueAST.Value))
		default:
			return nil
		}
	},
})
