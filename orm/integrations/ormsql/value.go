package ormsql

import (
	"fmt"
	"reflect"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/datatypes"
)

type valueCodec interface {
	goType() reflect.Type
	encode(protoValue protoreflect.Value, goValue reflect.Value)
}

var (
	timestampDesc = (&timestamppb.Timestamp{}).ProtoReflect().Descriptor()
	durationDesc  = (&durationpb.Duration{}).ProtoReflect().Descriptor()
)

func (b *builder) getValueCodec(descriptor protoreflect.FieldDescriptor) (valueCodec, error) {
	if descriptor.IsList() {
		return listCodec{jsonMarshalOptions: b.jsonMarshalOptions}, nil
	}

	if descriptor.IsMap() {
		return mapCodec{jsonMarshalOptions: b.jsonMarshalOptions}, nil
	}

	switch descriptor.Kind() {
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return int32Codec{}, nil
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return uint32Codec{}, nil
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return nil, fmt.Errorf("TODO")
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return nil, fmt.Errorf("TODO")
	case protoreflect.BoolKind:
		return nil, fmt.Errorf("TODO")
	case protoreflect.EnumKind:
		return nil, fmt.Errorf("TODO")
	case protoreflect.FloatKind:
		return nil, fmt.Errorf("TODO")
	case protoreflect.DoubleKind:
		return nil, fmt.Errorf("TODO")
	case protoreflect.BytesKind:
		return nil, fmt.Errorf("TODO")
	case protoreflect.StringKind:
		return nil, fmt.Errorf("TODO")
	case protoreflect.MessageKind:
		switch descriptor.Message().FullName() {
		case timestampDesc.FullName():
			return nil, fmt.Errorf("TODO")
		case durationDesc.FullName():
			return nil, fmt.Errorf("TODO")
		default:
			return messageValueCodec{jsonMarshalOptions: b.jsonMarshalOptions}, nil
		}
	default:
		panic("TODO")
	}
}

type int32Codec struct{}

func (i int32Codec) goType() reflect.Type { return reflect.TypeOf(int32(0)) }

func (i int32Codec) encode(protoValue protoreflect.Value, goValue reflect.Value) {
	goValue.SetInt(protoValue.Int())
}

type uint32Codec struct{}

func (u uint32Codec) goType() reflect.Type { return reflect.TypeOf(uint32(0)) }

func (u uint32Codec) encode(protoValue protoreflect.Value, goValue reflect.Value) {
	goValue.SetUint(protoValue.Uint())
}

type messageValueCodec struct {
	jsonMarshalOptions protojson.MarshalOptions
}

func (m messageValueCodec) goType() reflect.Type {
	return reflect.TypeOf(datatypes.JSON{})
}

func (m messageValueCodec) encode(protoValue protoreflect.Value, goValue reflect.Value) {
	bz, err := m.jsonMarshalOptions.Marshal(protoValue.Message().Interface())
	if err != nil {
		return
	}

	goValue.Set(reflect.ValueOf(datatypes.JSON(bz)))
}

type listCodec struct {
	jsonMarshalOptions protojson.MarshalOptions
}

func (l listCodec) goType() reflect.Type {
	return reflect.TypeOf(datatypes.JSON{})
}

func (l listCodec) encode(protoValue protoreflect.Value, goValue reflect.Value) {
	list := protoValue.List()
	n := list.Len()
	values := make([]interface{}, n)
	for i := 0; i < n; i++ {
		values[i] = list.Get(i).Interface()
	}
	structList, err := structpb.NewList(values)
	if err != nil {
		return
	}

	bz, err := l.jsonMarshalOptions.Marshal(structList)
	if err != nil {
		return
	}

	goValue.Set(reflect.ValueOf(datatypes.JSON(bz)))
}

type mapCodec struct {
	jsonMarshalOptions protojson.MarshalOptions
}

func (m mapCodec) goType() reflect.Type {
	return reflect.TypeOf(datatypes.JSON{})
}

func (m mapCodec) encode(protoValue protoreflect.Value, goValue reflect.Value) {
	protoMap := protoValue.Map()
	goMap := map[string]interface{}{}
	protoMap.Range(func(key protoreflect.MapKey, value protoreflect.Value) bool {
		goMap[key.String()] = value.Interface()
		return true
	})

	protoStruct, err := structpb.NewStruct(goMap)
	if err != nil {
		return
	}

	bz, err := m.jsonMarshalOptions.Marshal(protoStruct)
	if err != nil {
		return
	}

	goValue.Set(reflect.ValueOf(datatypes.JSON(bz)))
}
