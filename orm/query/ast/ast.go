package ast

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Expr interface {
	Eval(protoreflect.Value) protoreflect.Value
}

type Eq struct{ LHS, RHS Expr }

type LT struct{ LHS, RHS Expr }

type GT struct{ LHS, RHS Expr }

type LTE struct{ LHS, RHS Expr }

type GTE struct{ LHS, RHS Expr }

type And struct{ LHS, RHS Expr }

type Or struct{ LHS, RHS Expr }

type OrderBy struct {
	Desc  bool
	Field protoreflect.Name
}

type Query struct {
	Where         Expr
	Limit, Offset int
	OrderBy       []OrderBy
}
