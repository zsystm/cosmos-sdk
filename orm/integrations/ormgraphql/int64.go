package ormgraphql

import (
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type int64Codec struct{}

func (c int64Codec) Type() graphql.Type {
	return int64Scalar
}

func (c int64Codec) ToGraphql(value protoreflect.Value) (interface{}, error) {
	return value.Int(), nil
}

func (c int64Codec) FromGraphql(i interface{}) (protoreflect.Value, error) {
	return protoreflect.ValueOfInt64(i.(int64)), nil
}

var int64Scalar = graphql.NewScalar(graphql.ScalarConfig{
	Name: "int64",
	Serialize: func(value interface{}) interface{} {
		return strconv.FormatInt(value.(int64), 10)

	},
	ParseValue: func(value interface{}) interface{} {
		x, err := strconv.ParseInt(value.(string), 10, 64)
		if err != nil {
			return int64(0)
		}
		return x
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			x, err := strconv.ParseInt(valueAST.Value, 10, 64)
			if err != nil {
				return int64(0)
			}
			return x
		default:
			return nil
		}
	},
})
