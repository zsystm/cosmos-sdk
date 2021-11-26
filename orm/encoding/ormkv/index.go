package ormkv

import (
	"bytes"
	"io"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type IndexKeyCodec struct {
	*KeyCodec
	tableName       protoreflect.FullName
	indexFieldNames Fields
	pkFieldOrder    []int
	indexFields     []protoreflect.FieldDescriptor
}

func (cdc IndexKeyCodec) DecodeIndexKey(k, _ []byte) (indexFields, primaryKey []protoreflect.Value, err error) {

	values, err := cdc.Decode(bytes.NewReader(k))
	// got prefix key
	if err == io.EOF {
		return values, nil, nil
	} else if err != nil {
		return nil, nil, err
	}

	// got prefix key
	if len(values) < len(cdc.FieldCodecs) {
		return values, nil, nil
	}

	numPkFields := len(cdc.pkFieldOrder)
	pkValues := make([]protoreflect.Value, numPkFields)

	for i := 0; i < numPkFields; i++ {
		pkValues[i] = values[cdc.pkFieldOrder[i]]
	}

	return values, pkValues, nil
}

var _ IndexCodec = &IndexKeyCodec{}

func (cdc IndexKeyCodec) DecodeKV(k, v []byte) (Entry, error) {
	idxValues, pk, err := cdc.DecodeIndexKey(k, v)
	if err != nil {
		return nil, err
	}

	return &IndexKeyEntry{
		TableName:   cdc.tableName,
		Fields:      cdc.indexFieldNames,
		IndexValues: idxValues,
		PrimaryKey:  pk,
	}, nil
}

func (i IndexKeyCodec) EncodeKV(entry Entry) (k, v []byte, err error) {
	indexEntry, ok := entry.(*IndexKeyEntry)
	if !ok {
		return nil, nil, ormerrors.BadDecodeEntry
	}

	if indexEntry.TableName != i.tableName {
		return nil, nil, ormerrors.BadDecodeEntry
	}

	bz, err := i.KeyCodec.Encode(indexEntry.IndexValues)
	if err != nil {
		return nil, nil, err
	}

	return bz, sentinel, nil
}

var sentinel = []byte{0}

func (cdc IndexKeyCodec) ReadPrimaryKey(reader *bytes.Reader) ([]protoreflect.Value, error) {
	values, err := cdc.Decode(reader)
	if err != nil {
		return nil, err
	}

	return cdc.extractPrimaryKey(values), nil
}

func (cdc IndexKeyCodec) extractPrimaryKey(values []protoreflect.Value) []protoreflect.Value {
	numPkFields := len(cdc.pkFieldOrder)
	pkValues := make([]protoreflect.Value, numPkFields)

	for i := 0; i < numPkFields; i++ {
		pkValues[i] = values[cdc.pkFieldOrder[i]]
	}

	return pkValues
}

var _ Codec = &IndexKeyCodec{}

var _ Codec = &IndexKeyCodec{}
