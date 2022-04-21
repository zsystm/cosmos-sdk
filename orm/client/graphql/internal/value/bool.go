package value

import (
	"github.com/graphql-go/graphql"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type BoolCodec struct{}

func (b BoolCodec) Type() graphql.Type {
	return graphql.Boolean
}

func (b BoolCodec) ToGraphql(value protoreflect.Value) (interface{}, error) {
	return value.Bool(), nil
}

func (b BoolCodec) FromGraphql(i interface{}) (protoreflect.Value, error) {
	return protoreflect.ValueOfBool(i.(bool)), nil
}
