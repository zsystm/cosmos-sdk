package ast

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Var struct {
	Name string
}

func (v Var) Eval(c Context) protoreflect.Value {
	//TODO implement me
	panic("implement me")
}

func (v Var) Type() (*Type, error) {
	//TODO implement me
	panic("implement me")
}
