package ast

import "google.golang.org/protobuf/reflect/protoreflect"

type Context interface {
	Root() protoreflect.Value
	GetValue(name string) protoreflect.Value
}

type context struct {
	root protoreflect.Value
}

func (c context) GetValue(string) protoreflect.Value {
	return protoreflect.Value{}
}

func (c context) Root() protoreflect.Value {
	return c.root
}

func NewContextWithRoot(root protoreflect.Value) Context {
	return &context{root: root}
}
