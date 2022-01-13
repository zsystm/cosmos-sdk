package ormschema

import "google.golang.org/protobuf/reflect/protoreflect"

type ModuleDescriptor struct {
	FileDescriptors map[uint32]protoreflect.FileDescriptor
	Prefix          []byte
}
