package ormgraphql

import (
	"strings"

	"github.com/graphql-go/graphql"
	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/orm/encoding/ormkv"
)

func fieldsCamelCase(fields ormkv.Fields) {
	strcase.ToCamel(fields.String())
}

type fieldCodec interface {
	Type() graphql.Type
	ToGraphql(value protoreflect.Value) (interface{}, error)
	FromGraphql(interface{}) (protoreflect.Value, error)
}

func (b *Builder) protoMessageToGraphqlObject(descriptor protoreflect.MessageDescriptor) (*graphql.Object, error) {
	fullName := descriptor.FullName()

	obj, ok := b.protoObjects[fullName]
	if ok {
		return obj, nil
	}

	name := strings.ReplaceAll(string(fullName), ".", "_")
	fields, err := b.protoMessageToGraphqlFields(descriptor)
	if err != nil {
		return nil, err
	}

	obj = graphql.NewObject(graphql.ObjectConfig{
		Name:   name,
		Fields: fields,
		IsTypeOf: func(p graphql.IsTypeOfParams) bool {
			msg, ok := p.Value.(protoreflect.Message)
			if !ok {
				return false
			}
			return msg.Descriptor().FullName() == fullName
		},
		Description: getDocComments(descriptor),
	})
	b.protoObjects[fullName] = obj
	return obj, nil
}

func (b *Builder) protoMessageToGraphqlFields(descriptor protoreflect.MessageDescriptor) (graphql.Fields, error) {
	graphqlFields := map[string]*graphql.Field{}

	protoFields := descriptor.Fields()
	n := protoFields.Len()
	for i := 0; i < n; i++ {
		pf := protoFields.Get(i)
		gf, err := b.protoFieldToGraphqlField(pf)
		if err != nil {
			return nil, err
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

func (b *Builder) protoFieldToGraphqlField(fieldDescriptor protoreflect.FieldDescriptor) (*graphql.Field, error) {
	cdc, err := b.getFieldCodec(fieldDescriptor)
	if err != nil {
		return nil, err
	}

	resolver := func(p graphql.ResolveParams) (interface{}, error) {
		return cdc.ToGraphql(p.Source.(protoreflect.Message).Get(fieldDescriptor))
	}

	var deprecationReason string
	deprecated := protodesc.ToFieldDescriptorProto(fieldDescriptor).Options.Deprecated
	if deprecated != nil && *deprecated == true {
		deprecationReason = "field is marked as deprecated"
	}

	return &graphql.Field{
		Name:              fieldDescriptor.JSONName(),
		Type:              cdc.Type(),
		Resolve:           resolver,
		Subscribe:         resolver,
		DeprecationReason: deprecationReason,
		Description:       getDocComments(fieldDescriptor),
	}, nil
}

func (b *Builder) getFieldCodec(fieldDescriptor protoreflect.FieldDescriptor) (fieldCodec, error) {
	// TODO oneofs, maps

	cdc, err := getFieldCodecBasic(fieldDescriptor)
	if err != nil {
		return nil, err
	}

	if fieldDescriptor.IsList() {
		return listCodec{cdc}, nil
	} else {
		return cdc, nil
	}
}

func getFieldCodecBasic(fieldDescriptor protoreflect.FieldDescriptor) (fieldCodec, error) {
	switch fieldDescriptor.Kind() {
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return int32Codec{}, nil
	default:
		panic("TODO")
	}
}

type int32Codec struct{}

func (i int32Codec) FromGraphql(i2 interface{}) (protoreflect.Value, error) {
	return protoreflect.ValueOfInt32(i2.(int32)), nil
}

func (i int32Codec) Type() graphql.Type {
	return graphql.Int
}

func (i int32Codec) ToGraphql(value protoreflect.Value) (interface{}, error) {
	return int32(value.Int()), nil
}

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
