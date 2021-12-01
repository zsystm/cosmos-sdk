package ormgraphql

import (
	"github.com/graphql-go/graphql"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type stringCodec struct {
}

func (s stringCodec) Type() graphql.Type {
	return graphql.String
}

func (s stringCodec) ToGraphql(value protoreflect.Value) (interface{}, error) {
	return value.Interface().(string), nil
}

func (s stringCodec) FromGraphql(i interface{}) (protoreflect.Value, error) {
	return protoreflect.ValueOfString(i.(string)), nil
}
