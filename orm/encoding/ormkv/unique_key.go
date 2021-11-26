package ormkv

import (
	"bytes"
	"io"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type UniqueKeyCodec struct {
	tableName       protoreflect.FullName
	indexFieldNames Fields
	pkFieldOrder    []struct {
		inKey bool
		i     int
	}
	KeyCodec   *KeyCodec
	ValueCodec *KeyCodec
}

func (u UniqueKeyCodec) DecodeIndexKey(k, v []byte) (indexFields, primaryKey []protoreflect.Value, err error) {
	ks, err := u.KeyCodec.Decode(bytes.NewReader(k))

	// got prefix key
	if err == io.EOF {
		return ks, nil, err
	} else if err != nil {
		return nil, nil, err
	}

	// got prefix key
	if len(ks) < len(u.KeyCodec.FieldCodecs) {
		return ks, nil, err
	}

	vs, err := u.ValueCodec.Decode(bytes.NewReader(v))
	if err != nil {
		return nil, nil, err
	}

	pk := u.extractPrimaryKey(ks, vs)
	return ks, pk, nil
}

var _ IndexCodec = UniqueKeyCodec{}

//func NewUniqueKeyCodec(
//	prefix []byte,
//	messageType protoreflect.MessageType,
//	indexFields []protoreflect.FieldDescriptor,
//	primaryKeyFields []protoreflect.FieldDescriptor,
//) (*UniqueKeyCodec, error) {
//	keyCdc, err := NewKeyCodec(prefix, indexFields)
//	if err != nil {
//		return nil, err
//	}
//	valueCdc, err := NewKeyCodec(nil, primaryKeyFields)
//	return &UniqueKeyCodec{
//		MessageType: messageType,
//		KeyCodec:    keyCdc,
//		ValueCodec:  valueCdc,
//	}, err
//}

func (u UniqueKeyCodec) ExtractPrimaryKey(k, v []byte) ([]protoreflect.Value, error) {
	ks, err := u.KeyCodec.Decode(bytes.NewReader(k))
	if err != nil {
		return nil, err
	}

	vs, err := u.ValueCodec.Decode(bytes.NewReader(v))
	if err != nil {
		return nil, err
	}

	return u.extractPrimaryKey(ks, vs), nil
}

func (u UniqueKeyCodec) DecodeKV(k, v []byte) (Entry, error) {
	idxVals, pk, err := u.DecodeIndexKey(k, v)
	if err != nil {
		return nil, err
	}
	return IndexKeyEntry{
		TableName:   u.tableName,
		Fields:      u.indexFieldNames,
		IsUnique:    true,
		IndexValues: idxVals,
		PrimaryKey:  pk,
	}, err
}

func (cdc UniqueKeyCodec) extractPrimaryKey(keyValues, valueValues []protoreflect.Value) []protoreflect.Value {
	numPkFields := len(cdc.pkFieldOrder)
	pkValues := make([]protoreflect.Value, numPkFields)

	for i := 0; i < numPkFields; i++ {
		fo := cdc.pkFieldOrder[i]
		if fo.inKey {
			pkValues[i] = keyValues[fo.i]
		} else {
			pkValues[i] = valueValues[fo.i]
		}
	}

	return pkValues
}

func (u UniqueKeyCodec) EncodeKV(entry Entry) (k, v []byte, err error) {
	indexEntry := entry.(IndexKeyEntry)
	k, err = u.KeyCodec.Encode(indexEntry.IndexValues)
	if err != nil {
		return nil, nil, err
	}
	v, err = u.ValueCodec.Encode(indexEntry.PrimaryKey)
	return k, v, err
}

var _ Codec = &UniqueKeyCodec{}

//func NewUniqueKeyCodec(
//	messageDescriptor protoreflect.MessageDescriptor,
//	descriptor *ormpb.TableDescriptor,
//	id int,
//) {
//
//}
//
//func (cdc UniqueKeyCodec) DecodeKV(k, v []byte) (ormdecode.Entry, error) {
//	panic("TODO")
//}
