package ormkv

import (
	"bytes"
	"io"
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/orm/encoding/ormfield"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
)

type KeyCodec struct {
	prefix         []byte
	fixedSize      int
	variableSizers []struct {
		cdc ormfield.Codec
		i   int
	}
	FieldDescriptors []protoreflect.FieldDescriptor
	FieldCodecs      []ormfield.Codec
	FieldNames       []protoreflect.Name
}

var _ KeyCodec

func NewBaseCodec(prefix []byte, fieldDescs []protoreflect.FieldDescriptor) (*KeyCodec, error) {
	n := len(fieldDescs)
	var valueCodecs []ormfield.Codec
	var variableSizers []struct {
		cdc ormfield.Codec
		i   int
	}
	fixedSize := 0
	names := make([]protoreflect.Name, len(fieldDescs))
	for i := 0; i < n; i++ {
		nonTerminal := true
		if i == n-1 {
			nonTerminal = false
		}
		field := fieldDescs[i]
		cdc, err := ormfield.GetCodec(field, nonTerminal)
		if err != nil {
			return nil, err
		}
		if x := cdc.FixedBufferSize(); x > 0 {
			fixedSize += x
		} else {
			variableSizers = append(variableSizers, struct {
				cdc ormfield.Codec
				i   int
			}{cdc, i})
		}
		valueCodecs = append(valueCodecs, cdc)
		names[i] = field.Name()
	}

	return &KeyCodec{
		FieldCodecs:      valueCodecs,
		FieldDescriptors: fieldDescs,
		FieldNames:       names,
		prefix:           prefix,
		fixedSize:        fixedSize,
		variableSizers:   variableSizers,
	}, nil
}

func (cdc *KeyCodec) Encode(values []protoreflect.Value) ([]byte, error) {
	sz, err := cdc.ComputeBufferSize(values)
	if err != nil {
		return nil, err
	}

	w := bytes.NewBuffer(make([]byte, 0, sz))
	_, err = w.Write(cdc.prefix)
	if err != nil {
		return nil, err
	}

	n := len(values)
	if n > len(cdc.FieldCodecs) {
		return nil, ormerrors.IndexOutOfBounds
	}

	for i := 0; i < n; i++ {
		err = cdc.FieldCodecs[i].Encode(values[i], w)
		if err != nil {
			return nil, err
		}
	}
	return w.Bytes(), nil
}

func (cdc *KeyCodec) GetValues(message protoreflect.Message) []protoreflect.Value {
	var res []protoreflect.Value
	for _, f := range cdc.FieldDescriptors {
		res = append(res, message.Get(f))
	}
	return res
}

func SkipPrefix(r *bytes.Reader, prefix []byte) error {
	n := len(prefix)
	if n > 0 {
		// we skip checking the prefix for performance reasons because we assume
		// that it was checked by the caller
		_, err := r.Seek(int64(n), io.SeekCurrent)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cdc *KeyCodec) Decode(r *bytes.Reader) ([]protoreflect.Value, error) {
	err := SkipPrefix(r, cdc.prefix)
	if err != nil {
		return nil, err
	}

	n := len(cdc.FieldCodecs)
	values := make([]protoreflect.Value, n)
	for i := 0; i < n; i++ {
		value, err := cdc.FieldCodecs[i].Decode(r)
		values[i] = value
		if err == io.EOF {
			if i == n-1 {
				return values, nil
			} else {
				return nil, io.ErrUnexpectedEOF
			}
		} else if err != nil {
			return nil, err
		}
	}
	return values, nil
}

func (cdc *KeyCodec) EncodeFromMessage(message protoreflect.Message) ([]protoreflect.Value, []byte, error) {
	values := cdc.GetValues(message)
	bz, err := cdc.Encode(values)
	return values, bz, err
}

// IsFullyOrdered returns true if all parts are also ordered
func (cdc *KeyCodec) IsFullyOrdered() bool {
	for _, p := range cdc.FieldCodecs {
		if !p.IsOrdered() {
			return false
		}
	}
	return true
}

func (cdc *KeyCodec) CompareValues(values1, values2 []protoreflect.Value) int {
	n := len(values1)
	if n != len(values2) {
		panic("expected arrays of the same length")
	}
	if n > len(cdc.FieldCodecs) {
		panic("array is too long")
	}

	var cmp int
	for i := 0; i < n; i++ {
		cmp = cdc.FieldCodecs[i].Compare(values1[i], values2[i])
		// any non-equal parts determine our ordering
		if cmp != 0 {
			break
		}
	}

	return cmp
}

func GetFieldDescriptors(desc protoreflect.MessageDescriptor, fields string) ([]protoreflect.FieldDescriptor, error) {
	if len(fields) == 0 {
		return nil, ormerrors.InvalidKeyFieldsDefinition.Wrapf("got fields %q for table %q", fields, desc.FullName())
	}

	fieldNames := strings.Split(fields, ",")

	have := map[string]bool{}

	var fieldDescriptors []protoreflect.FieldDescriptor
	for _, fieldName := range fieldNames {
		if have[fieldName] {
			return nil, ormerrors.DuplicateKeyField.Wrapf("field %q in %q", fieldName, desc.FullName())
		}

		have[fieldName] = true
		fieldDesc := desc.Fields().ByName(protoreflect.Name(fieldName))
		if fieldDesc == nil {
			return nil, ormerrors.FieldNotFound.Wrapf("field %q in %q", fieldName, desc.FullName())
		}

		fieldDescriptors = append(fieldDescriptors, fieldDesc)
	}
	return fieldDescriptors, nil
}

func (cdc KeyCodec) ComputeBufferSize(values []protoreflect.Value) (int, error) {
	size := cdc.fixedSize
	n := len(values)
	for _, sz := range cdc.variableSizers {
		if sz.i >= n {
			return size, nil
		}
		x, err := sz.cdc.ComputeBufferSize(values[sz.i])
		if err != nil {
			return 0, err
		}
		size += x
	}
	return size, nil
}

func (cdc *KeyCodec) SetValues(message protoreflect.Message, values []protoreflect.Value) {
	for i, f := range cdc.FieldDescriptors {
		message.Set(f, values[i])
	}
}

func (cdc KeyCodec) Prefix() []byte {
	return cdc.prefix
}

// CheckValidRangeIterationKeys checks if the start and end key prefixes are valid
// for range iteration meaning that for each non-equal field in the prefixes
// those field types support ordered iteration.
func (cdc KeyCodec) CheckValidRangeIterationKeys(start, end []protoreflect.Value) error {
	n := len(start)
	if len(end) < n {
		n = len(end)
	}

	for i := 0; i < n; i++ {
		fieldCdc := cdc.FieldCodecs[i]
		if fieldCdc.IsOrdered() {
			continue
		}
		x := start[i]
		y := end[i]
		if fieldCdc.Compare(x, y) != 0 {
			descriptor := cdc.FieldDescriptors[i]
			return ormerrors.InvalidRangeIterationKeys.Wrapf(

				"field %s of kind %s doesn't support ordered range iteration",
				descriptor.FullName(),
				descriptor.Kind(),
			)
		}
	}
	return nil
}
