package graphql

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/orm/client/graphql/internal/value"
)

func (b *Builder) protoFieldToGraphqlField(fieldDescriptor protoreflect.FieldDescriptor) (*graphql.Field, error) {
	cdc, err := b.getFieldCodec(fieldDescriptor)
	if err != nil {
		return nil, err
	}

	resolver := func(p graphql.ResolveParams) (interface{}, error) {
		return cdc.ToGraphql(p.Source.(protoreflect.Message).Get(fieldDescriptor))
	}

	var deprecationReason string
	fdProtoOpts := protodesc.ToFieldDescriptorProto(fieldDescriptor).Options
	if fdProtoOpts != nil {
		deprecated := fdProtoOpts.Deprecated
		if deprecated != nil && *deprecated == true {
			deprecationReason = "field is marked as deprecated"
		}
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

func (b *Builder) getFieldCodec(fieldDescriptor protoreflect.FieldDescriptor) (value.ValueCodec, error) {
	// TODO maps

	cdc, err := b.getFieldCodecBasic(fieldDescriptor)
	if err != nil {
		return nil, err
	}

	if fieldDescriptor.IsList() {
		return value.ListCodec{BasicCodec: cdc}, nil
	} else {
		return cdc, nil
	}
}

func (b *Builder) getFieldCodecBasic(fieldDescriptor protoreflect.FieldDescriptor) (value.ValueCodec, error) {
	switch fieldDescriptor.Kind() {
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return value.Int32Codec{}, nil
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return value.Int64Codec{}, nil
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return value.Uint32Codec{}, nil
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return value.Uint64Codec{}, nil
	case protoreflect.BoolKind:
		return value.BoolCodec{}, nil
	case protoreflect.StringKind:
		return value.StringCodec{}, nil
	case protoreflect.EnumKind:
		enum, err := b.protoEnumToGraphqlEnum(fieldDescriptor.Enum())
		if err != nil {
			return nil, err
		}
		return value.EnumCodec{Enum: enum}, nil
	case protoreflect.MessageKind:
		if fieldDescriptor.IsMap() {
			return nil, fmt.Errorf("maps not supported yet")
		}
		obj, err := b.protoMessageToGraphqlObject(fieldDescriptor.Message())
		if err != nil {
			return nil, err
		}
		return messageCodec{obj}, nil
	case protoreflect.BytesKind:
		return value.BytesCodec{}, nil
	default:
		return nil, fmt.Errorf("field of kind %v not supported", fieldDescriptor.Kind())
	}
}

type messageCodec struct {
	obj *graphql.Object
}

func (m messageCodec) Type() graphql.Type {
	return m.obj
}

func (m messageCodec) ToGraphql(value protoreflect.Value) (interface{}, error) {
	return value.Message(), nil
}

func (m messageCodec) FromGraphql(i interface{}) (protoreflect.Value, error) {
	return protoreflect.ValueOfMessage(i.(protoreflect.Message)), nil
}
