package value

import (
	"github.com/graphql-go/graphql"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type ValueCodec interface {
	Type() graphql.Type
	ToGraphql(value protoreflect.Value) (interface{}, error)
	FromGraphql(interface{}) (protoreflect.Value, error)
}
