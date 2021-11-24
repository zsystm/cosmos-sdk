package ormkv

import (
	"bytes"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type KeyCodec interface {
	// Encode encodes the values assuming that they correspond to the fields
	// specified for the key. If the array of values is shorter than the
	// number of fields in the key, a partial "prefix" key will be encoded
	// which can be used for constructing a prefix iterator.
	Encode(values []protoreflect.Value) ([]byte, error)

	// Decode decodes the values in the key specified by the reader. If the
	// provided key is a prefix key, the values that could be decoded will
	// be returned with io.EOF as the error.
	Decode(r *bytes.Reader) ([]protoreflect.Value, error)

	// GetValues extracts the values specified by the key fields from the message.
	GetValues(mref protoreflect.Message) []protoreflect.Value

	// EncodeFromMessage combines GetValues and Encode.
	EncodeFromMessage(message protoreflect.Message) ([]protoreflect.Value, []byte, error)

	IsFullyOrdered() bool

	CompareValues(values1, values2 []protoreflect.Value) int

	ComputeBufferSize(values []protoreflect.Value) (int, error)
	Prefix() []byte
}

type Codec interface {
	DecodeKV(k, v []byte) (Entry, error)
	EncodeKV(entry Entry) (k, v []byte, err error)
}
