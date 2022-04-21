package value

import (
	"github.com/graphql-go/graphql"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type StringCodec struct {
}

func (s StringCodec) Type() graphql.Type {
	return graphql.String
}

func (s StringCodec) ToGraphql(value protoreflect.Value) (interface{}, error) {
	return value.Interface().(string), nil
}

func (s StringCodec) FromGraphql(i interface{}) (protoreflect.Value, error) {
	return protoreflect.ValueOfString(i.(string)), nil
}
