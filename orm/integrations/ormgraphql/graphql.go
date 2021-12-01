package ormgraphql

import (
	"github.com/graphql-go/graphql"
)

type Builder struct {
	objects map[string]*graphql.Object
	enums   map[string]*graphql.Enum
}

func NewBuilder() *Builder {
	return &Builder{
		objects: map[string]*graphql.Object{},
		enums:   map[string]*graphql.Enum{},
	}
}

//func (b Builder) buildTable(tableDesc *ormpb.TableDescriptor, desc protoreflect.MessageDescriptor) (*graphql.Field, error) {
//	name := descriptorName(desc)
//	objType, err := b.protoMessageToGraphqlObject(desc)
//	if err != nil {
//		return nil, err
//	}
//
//	return &graphql.Field{
//		Name:              name,
//		Type:              objType,
//		Args:              nil,
//		Resolve:           nil,
//		DeprecationReason: "",
//		Description:       getDocComments(desc),
//	}, nil
//}
