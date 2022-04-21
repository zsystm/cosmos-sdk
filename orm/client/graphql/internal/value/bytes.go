package value

import (
	"encoding/base64"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type BytesCodec struct{}

func (b BytesCodec) Type() graphql.Type {
	return bytesScalar
}

func (b BytesCodec) ToGraphql(value protoreflect.Value) (interface{}, error) {
	return value.Bytes(), nil
}

func (b BytesCodec) FromGraphql(i interface{}) (protoreflect.Value, error) {
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
			return DecodeBase64Scalar(value)
		case *string:
			return DecodeBase64Scalar(*value)
		default:
			return nil
		}
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			return DecodeBase64Scalar(valueAST.Value)
		default:
			return nil
		}
	},
})
