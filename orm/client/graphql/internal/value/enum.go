package value

import (
	"github.com/graphql-go/graphql"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type EnumCodec struct {
	Enum *graphql.Enum
}

func (e EnumCodec) Type() graphql.Type {
	return e.Enum
}

func (e EnumCodec) ToGraphql(value protoreflect.Value) (interface{}, error) {
	return value.Enum(), nil
}

func (e EnumCodec) FromGraphql(i interface{}) (protoreflect.Value, error) {
	return protoreflect.ValueOfEnum(i.(protoreflect.EnumNumber)), nil
}
