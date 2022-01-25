package ast

import "google.golang.org/protobuf/reflect/protoreflect"

type Field struct {
	Target Expr
	Field  protoreflect.FieldDescriptor
}

func (f *Field) Eval(ctx Context) protoreflect.Value {
	return f.Target.Eval(ctx).Message().Get(f.Field)
}

func (f *Field) Type() (*Type, error) {
	return &Type{
		Kind:     f.Field.Kind(),
		Message:  f.Field.Message(),
		Repeated: f.Field.Cardinality() == protoreflect.Repeated,
	}, nil
}
