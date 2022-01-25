package ast

import "google.golang.org/protobuf/reflect/protoreflect"

type Root struct {
	MessageDescriptor protoreflect.MessageDescriptor
}

func (r Root) Eval(ctx Context) protoreflect.Value {
	return ctx.Root()
}

func (r Root) Type() (*Type, error) {
	return &Type{
		Kind:     protoreflect.MessageKind,
		Message:  r.MessageDescriptor,
		Repeated: false,
	}, nil
}
