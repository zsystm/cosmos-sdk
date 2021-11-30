package ormgraphql

import (
	"github.com/graphql-go/graphql"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type boolCodec struct{}

func (b boolCodec) Type() graphql.Type {
	return graphql.Boolean
}

func (b boolCodec) ToGraphql(value protoreflect.Value) (interface{}, error) {
	return value.Bool(), nil
}

func (b boolCodec) FromGraphql(i interface{}) (protoreflect.Value, error) {
	return protoreflect.ValueOfBool(i.(bool)), nil
}
