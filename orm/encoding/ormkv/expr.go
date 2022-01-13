package ormkv

import (
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type Expression interface {
	fmt.Stringer

	GetValue(protoreflect.Message) protoreflect.Value
}

type FieldExpression struct {
	protoreflect.FieldDescriptor
}

func (f FieldExpression) String() string {
	return string(f.Name())
}

func (f FieldExpression) GetValue(message protoreflect.Message) protoreflect.Value {
	return message.Get(f.FieldDescriptor)
}

var _ Expression = FieldExpression{}
