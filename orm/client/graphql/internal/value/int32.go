package value

import (
	"github.com/graphql-go/graphql"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Int32Codec struct{}

func (i Int32Codec) FromGraphql(i2 interface{}) (protoreflect.Value, error) {
	return protoreflect.ValueOfInt32(i2.(int32)), nil
}

func (i Int32Codec) Type() graphql.Type {
	return graphql.Int
}

func (i Int32Codec) ToGraphql(value protoreflect.Value) (interface{}, error) {
	return int32(value.Int()), nil
}
