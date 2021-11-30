package ormgraphql

import (
	"github.com/graphql-go/graphql"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type listCodec struct {
	basicCodec fieldCodec
}

func (l listCodec) Type() graphql.Type {
	return graphql.NewList(graphql.NewNonNull(l.basicCodec.Type()))
}

func (l listCodec) ToGraphql(value protoreflect.Value) (interface{}, error) {
	list := value.List()
	n := list.Len()
	res := make([]interface{}, n)
	var err error
	for i := 0; i < n; i++ {
		res[i], err = l.basicCodec.ToGraphql(list.Get(i))
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (l listCodec) FromGraphql(i interface{}) (protoreflect.Value, error) {
	panic("implement me")
}
