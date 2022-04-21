package value

import (
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Uint32Codec struct{}

func (u Uint32Codec) Type() graphql.Type {
	return uint32Scalar
}

func (u Uint32Codec) ToGraphql(value protoreflect.Value) (interface{}, error) {
	return uint32(value.Uint()), nil
}

func (u Uint32Codec) FromGraphql(i interface{}) (protoreflect.Value, error) {
	return protoreflect.ValueOfUint32(i.(uint32)), nil
}

var uint32Scalar = graphql.NewScalar(graphql.ScalarConfig{
	Name: "UInt32",
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
