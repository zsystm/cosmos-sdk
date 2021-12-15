package ormkv

import (
	"bytes"
	"io"

	"google.golang.org/protobuf/reflect/protoreflect"
)

// SkipPrefix skips the provided prefix in the reader or returns an error.
// This is used for efficient logical decoding of keys.
func SkipPrefix(r *bytes.Reader, prefix []byte) error {
	n := len(prefix)
	if n > 0 {
		// we skip checking the prefix for performance reasons because we assume
		// that it was checked by the caller
		_, err := r.Seek(int64(n), io.SeekCurrent)
		return err
	}
	return nil
}

func ValuesOf(values ...interface{}) []protoreflect.Value {
	n := len(values)
	res := make([]protoreflect.Value, n)
	for i := 0; i < n; i++ {
		res[i] = protoreflect.ValueOf(values[i])
	}
	return res
}
