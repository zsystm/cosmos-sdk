package ormsql

import (
	"fmt"
	"reflect"
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type fieldCodec struct {
	valueCodec
	structField reflect.StructField
	protoField  protoreflect.FieldDescriptor
}

func (b *builder) makeFieldCodec(descriptor protoreflect.FieldDescriptor, isPrimaryKey bool) (*fieldCodec, error) {
	valCdc, err := b.getValueCodec(descriptor)
	if err != nil {
		return nil, err
	}

	tag := fmt.Sprintf(`gorm:"column:%s`, descriptor.Name())
	if isPrimaryKey {
		tag = tag + fmt.Sprintf(`;primaryKey;autoIncrement:false`)
	}
	var fieldName = strings.ToTitle(string(descriptor.Name()))
	structField := reflect.StructField{
		Name: fieldName,
		Type: valCdc.goType(),
		Tag:  reflect.StructTag(tag + `"`),
	}

	return &fieldCodec{
		valueCodec:  valCdc,
		structField: structField,
		protoField:  descriptor,
	}, nil
}

func (f fieldCodec) encode(message protoreflect.Message, value reflect.Value) {
	if !message.Has(f.protoField) {
		return
	}

	protoVal := message.Get(f.protoField)
	goField := value.FieldByName(f.structField.Name)
	f.valueCodec.encode(protoVal, goField)
}
