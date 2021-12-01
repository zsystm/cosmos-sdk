package ormgraphql

import (
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type uint64Codec struct{}

func (u uint64Codec) Type() graphql.Type {
	return uint64Scalar
}

func (u uint64Codec) ToGraphql(value protoreflect.Value) (interface{}, error) {
	return value.Uint(), nil
}

func (u uint64Codec) FromGraphql(i interface{}) (protoreflect.Value, error) {
	return protoreflect.ValueOfUint64(i.(uint64)), nil
}

var uint64Scalar = graphql.NewScalar(graphql.ScalarConfig{
	Name: "uint64",
	Serialize: func(value interface{}) interface{} {
		return strconv.FormatUint(value.(uint64), 10)

	},
	ParseValue: func(value interface{}) interface{} {
		x, err := strconv.ParseUint(value.(string), 10, 64)
		if err != nil {
			return uint64(0)
		}
		return x
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			x, err := strconv.ParseUint(valueAST.Value, 10, 64)
			if err != nil {
				return uint64(0)
			}
			return x
		default:
			return nil
		}
	},
})
