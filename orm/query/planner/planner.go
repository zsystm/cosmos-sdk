package planner

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/query/ast"
	"github.com/cosmos/cosmos-sdk/orm/query/op"
)

func Plan(table ormtable.Table, query *ast.Query) (op.Op, error) {
	return naivePlan(table, query)
}

func naivePlan(table ormtable.Table, query *ast.Query) (op.Op, error) {
	var res op.Op
	res = &op.Filter{
		Func: func(message proto.Message) bool {
			ctx := ast.NewContextWithRoot(protoreflect.ValueOfMessage(message.ProtoReflect()))
			return query.Where.Eval(ctx).Bool()
		},
		Op: &op.IndexScan{
			Index:   table,
			Options: nil,
		},
	}

	for range query.OrderBy {
		panic("TODO")
	}

	return res, nil
}

func orderByToExpr(by ast.OrderBy) ast.Expr {
	x := &ast.Field{
		Target: &ast.Var{Name: "x"},
		Field:  by.Field,
	}
	y := &ast.Field{
		Target: &ast.Var{Name: "y"},
		Field:  by.Field,
	}
	if by.Desc {
		return &ast.GT{LHS: x, RHS: y}
	} else {
		return &ast.LT{LHS: x, RHS: y}
	}
}
