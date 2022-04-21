package graphql

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

func getDocComments(desc protoreflect.Descriptor) string {
	return desc.ParentFile().SourceLocations().ByDescriptor(desc).LeadingComments
}
