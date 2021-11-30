package ormgraphql

import (
	"encoding/base64"

	"google.golang.org/protobuf/reflect/protoreflect"
)

func decodeBase64Scalar(value string) []byte {
	bytes, err := base64.RawURLEncoding.DecodeString(value)
	if err != nil {
		return nil
	}
	return bytes
}

func getDocComments(desc protoreflect.Descriptor) string {
	return desc.ParentFile().SourceLocations().ByDescriptor(desc).LeadingComments
}
