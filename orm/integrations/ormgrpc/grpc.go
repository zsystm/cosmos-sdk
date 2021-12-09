package ormgrpc

import (
	"google.golang.org/protobuf/types/descriptorpb"
)

type fileBuilder struct {
}

type serviceBuilder struct {
	*fileBuilder
}

func (b *fileBuilder) addMessage(proto *descriptorpb.DescriptorProto) {

}

//func (b *builder) buildTableQueryMethods(
//	messageDescriptor protoreflect.MessageDescriptor,
//	tableDescriptor *ormv1alpha1.TableDescriptor,
//) error {
//
//}

//func (b *serviceBuilder) buildTableListMethod(messageDescriptor protoreflect.MessageDescriptor) error {
//	orderEnum := &descriptorpb.EnumDescriptorProto{
//		Name:  nil,
//		Value: nil,
//	}
//
//	inputType := &descriptorpb.DescriptorProto{
//		Name:           nil,
//		Field:          nil,
//		Extension:      nil,
//		NestedType:     nil,
//		EnumType:       nil,
//		ExtensionRange: nil,
//		OneofDecl:      nil,
//		Options:        nil,
//		ReservedRange:  nil,
//		ReservedName:   nil,
//	}
//	b.addMessage(inputType)
//
//	fields := messageDescriptor.Fields()
//	for i := 0; i < fields.Len(); i++ {
//		//field := fields.Get(i)
//		//enumValue := &descriptorpb.EnumValueDescriptorProto{
//		//	Name:    nil,
//		//	Number:  nil,
//		//	Options: nil,
//		//}
//		//orderEnum.Value = append(orderEnum.Value)
//	}
//
//	outputType := &descriptorpb.DescriptorProto{
//		Name:           nil,
//		Field:          nil,
//		Extension:      nil,
//		NestedType:     nil,
//		EnumType:       nil,
//		ExtensionRange: nil,
//		OneofDecl:      nil,
//		Options:        nil,
//		ReservedRange:  nil,
//		ReservedName:   nil,
//	}
//	b.addMessage(outputType)
//
//	method := &descriptorpb.MethodDescriptorProto{
//		Name:            nil,
//		InputType:       inputType.Name,
//		OutputType:      outputType.Name,
//		Options:         nil,
//		ClientStreaming: nil,
//		ServerStreaming: nil,
//	}
//
//	return nil
//}
