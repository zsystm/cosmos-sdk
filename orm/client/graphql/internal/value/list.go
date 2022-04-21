package value

import (
	"github.com/graphql-go/graphql"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type ListCodec struct {
	BasicCodec ValueCodec
}

func (l ListCodec) Type() graphql.Type {
	return graphql.NewList(graphql.NewNonNull(l.BasicCodec.Type()))
}

func (l ListCodec) ToGraphql(value protoreflect.Value) (interface{}, error) {
	list := value.List()
	n := list.Len()
	res := make([]interface{}, n)
	var err error
	for i := 0; i < n; i++ {
		res[i], err = l.BasicCodec.ToGraphql(list.Get(i))
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (l ListCodec) FromGraphql(i interface{}) (protoreflect.Value, error) {
	panic("implement me")
}
