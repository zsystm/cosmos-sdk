package ast

import "google.golang.org/protobuf/reflect/protoreflect"

type Expr interface {
	Eval(Context) protoreflect.Value
	Type() (*Type, error)
}

type Type struct {
	Kind     protoreflect.Kind
	Message  protoreflect.MessageDescriptor
	Repeated bool
}
