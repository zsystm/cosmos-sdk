package ormsql

import (
	"reflect"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/structpb"
	"gorm.io/datatypes"
)

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

func (m mapCodec) decode(goValue reflect.Value, protoValue protoreflect.Value) {
	//TODO implement me
	panic("implement me")
}
