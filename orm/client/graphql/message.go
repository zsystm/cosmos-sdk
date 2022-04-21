package graphql

import (
	"fmt"
	"strings"

	"github.com/graphql-go/graphql"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (b *Builder) registerMessages(descriptors protoreflect.MessageDescriptors) error {
	n := descriptors.Len()
	for i := 0; i < n; i++ {
		_, err := b.protoMessageToGraphqlObject(descriptors.Get(i))
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Builder) protoMessageToGraphqlObject(descriptor protoreflect.MessageDescriptor) (*graphql.Object, error) {
	name := descriptorName(descriptor)

	if obj, ok := b.objects[name]; ok {
		return obj, nil
	}

	fields, err := b.protoMessageToGraphqlFields(descriptor)
	if err != nil {
		return nil, err
	}

	obj := graphql.NewObject(graphql.ObjectConfig{
		Name:   name,
		Fields: fields,
	})
	b.objects[name] = obj

	err = b.registerMessages(descriptor.Messages())
	return obj, err
}

func descriptorName(descriptor protoreflect.Descriptor) string {
	return strings.ReplaceAll(string(descriptor.FullName()), ".", "_")
}

func (b *Builder) protoMessageToGraphqlFields(descriptor protoreflect.MessageDescriptor) (graphql.Fields, error) {
	graphqlFields := map[string]*graphql.Field{}

	protoFields := descriptor.Fields()
	n := protoFields.Len()
	for i := 0; i < n; i++ {
		pf := protoFields.Get(i)
		gf, err := b.protoFieldToGraphqlField(pf)
		if err != nil {
			//return nil, err
			fmt.Printf("error resolving field %v: %v\n", pf, err)
			continue
		}
		graphqlFields[gf.Name] = gf
	}

	protoOneofs := descriptor.Oneofs()
	n = protoOneofs.Len()
	for i := 0; i < n; i++ {
		//oneof := protoOneofs.Get(i)
	}

	return graphqlFields, nil
}
