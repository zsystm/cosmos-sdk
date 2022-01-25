package ast

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Eq struct{ LHS, RHS Expr }

type LT struct{ LHS, RHS Expr }

func (L LT) Eval(c Context) protoreflect.Value {
	//TODO implement me
	panic("implement me")
}

func (L LT) Type() (*Type, error) {
	//TODO implement me
	panic("implement me")
}

type GT struct{ LHS, RHS Expr }

func (G GT) Eval(c Context) protoreflect.Value {
	//TODO implement me
	panic("implement me")
}

func (G GT) Type() (*Type, error) {
	//TODO implement me
	panic("implement me")
}

type LTE struct{ LHS, RHS Expr }

type GTE struct{ LHS, RHS Expr }

type And struct{ Exprs []Expr }

type Or struct{ Exprs []Expr }

type OrderBy struct {
	Desc  bool
	Field protoreflect.FieldDescriptor
}

type Query struct {
	Where         Expr
	Limit, Offset int
	OrderBy       []OrderBy
}
