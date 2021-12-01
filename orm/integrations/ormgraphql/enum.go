package ormgraphql

import (
	"github.com/graphql-go/graphql"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (b *Builder) protoEnumToGraphqlEnum(descriptor protoreflect.EnumDescriptor) (*graphql.Enum, error) {
	name := descriptorName(descriptor)
	if enum, ok := b.enums[name]; ok {
		return enum, nil
	}

	enumValueConfigMap := map[string]*graphql.EnumValueConfig{}

	protoValues := descriptor.Values()
	n := protoValues.Len()
	for i := 0; i < n; i++ {
		protoValue := protoValues.Get(i)
		var deprecationReason string
		evdProtoOpts := protodesc.ToEnumValueDescriptorProto(protoValue).Options
		if evdProtoOpts != nil {
			deprecated := evdProtoOpts.Deprecated
			if deprecated != nil && *deprecated == true {
				deprecationReason = "enum value is marked as deprecated"
			}
		}

		enumValueConfigMap[string(protoValue.Name())] = &graphql.EnumValueConfig{
			Value:             protoValue.Number(),
			DeprecationReason: deprecationReason,
			Description:       getDocComments(protoValue),
		}
	}

	enum := graphql.NewEnum(graphql.EnumConfig{
		Name:        name,
		Values:      enumValueConfigMap,
		Description: getDocComments(descriptor),
	})

	b.enums[enum.Name()] = enum
	return enum, nil
}

type enumCodec struct {
	enum *graphql.Enum
}

func (e enumCodec) Type() graphql.Type {
	return e.enum
}

func (e enumCodec) ToGraphql(value protoreflect.Value) (interface{}, error) {
	return value.Enum(), nil
}

func (e enumCodec) FromGraphql(i interface{}) (protoreflect.Value, error) {
	return protoreflect.ValueOfEnum(i.(protoreflect.EnumNumber)), nil
}
