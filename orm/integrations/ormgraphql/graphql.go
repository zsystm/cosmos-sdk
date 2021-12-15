package ormgraphql

import (
	"github.com/graphql-go/graphql"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Builder struct {
	objects map[string]*graphql.Object
	enums   map[string]*graphql.Enum
	query   map[string]*graphql.Object
}

func NewBuilder() *Builder {
	return &Builder{
		objects: map[string]*graphql.Object{},
		enums:   map[string]*graphql.Enum{},
		query:   map[string]*graphql.Object{},
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

func (b *Builder) RegisterFileDescriptor(descriptor protoreflect.FileDescriptor) error {
	return b.registerMessages(descriptor.Messages())
}

func (b *Builder) Build() (graphql.Schema, error) {
	return graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:        "Query",
			Fields:      b.query,
			Description: "The root query object.",
		}),
	})
}

//func Execute(schema graphql.Schema, store kv.ReadMultiStore, requestString string) {
//	graphql.Do(graphql.Params{
//		Schema:         schema,
//		RequestString:  "",
//		RootObject:     nil,
//		VariableValues: nil,
//		OperationName:  "",
//		Context:        nil,
//	})
//}
