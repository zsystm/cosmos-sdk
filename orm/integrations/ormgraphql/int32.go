package ormgraphql

import (
	"github.com/graphql-go/graphql"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type int32Codec struct{}

func (i int32Codec) FromGraphql(i2 interface{}) (protoreflect.Value, error) {
	return protoreflect.ValueOfInt32(i2.(int32)), nil
}

func (i int32Codec) Type() graphql.Type {
	return graphql.Int
}

func (i int32Codec) ToGraphql(value protoreflect.Value) (interface{}, error) {
	return int32(value.Int()), nil
}
