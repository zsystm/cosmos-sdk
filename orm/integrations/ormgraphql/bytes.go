package ormgraphql

import (
	"encoding/base64"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type bytesCodec struct{}

func (b bytesCodec) Type() graphql.Type {
	return bytesScalar
}

func (b bytesCodec) ToGraphql(value protoreflect.Value) (interface{}, error) {
	return value.Bytes(), nil
}

func (b bytesCodec) FromGraphql(i interface{}) (protoreflect.Value, error) {
	return protoreflect.ValueOfBytes(i.([]byte)), nil
}

var bytesScalar = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "Cursor",
	Description: "",
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case []byte:
			return base64.RawURLEncoding.EncodeToString(value)
		default:
			return nil
		}
	},
	ParseValue: func(value interface{}) interface{} {
		switch value := value.(type) {
		case string:
			return decodeBase64Scalar(value)
		case *string:
			return decodeBase64Scalar(*value)
		default:
			return nil
		}
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			return decodeBase64Scalar(valueAST.Value)
		default:
			return nil
		}
	},
})
