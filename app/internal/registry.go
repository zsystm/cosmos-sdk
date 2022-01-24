package internal

import "google.golang.org/protobuf/reflect/protoreflect"

var Registry = map[protoreflect.FullName]*Initializer{}
var Declarations = map[protoreflect.FullName]map[interface{}]interface{}{}

type Initializer struct {
	Type              protoreflect.MessageType
	Providers         []interface{}
	ProviderFactories []func(ModuleDeclarations)
}

type ModuleDeclarations struct {
	Declarations map[protoreflect.FullName]map[interface{}]interface{}
}

type Option interface {
	apply(*Initializer)
}
