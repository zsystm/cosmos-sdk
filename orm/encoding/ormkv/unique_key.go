package ormkv

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

type UniqueKeyCodec struct {
	MessageType protoreflect.MessageType
	KeyCodec    *KeyCodec
	ValueCodec  *KeyCodec
}

func NewUniqueKeyCodec(
	prefix []byte,
	messageType protoreflect.MessageType,
	indexFields []protoreflect.FieldDescriptor,
	primaryKeyFields []protoreflect.FieldDescriptor,
) (*UniqueKeyCodec, error) {
	keyCdc, err := NewBaseCodec(prefix, indexFields)
	if err != nil {
		return nil, err
	}
	valueCdc, err := NewBaseCodec(nil, primaryKeyFields)
	return &UniqueKeyCodec{
		MessageType: messageType,
		KeyCodec:    keyCdc,
		ValueCodec:  valueCdc,
	}, err
}

func (u UniqueKeyCodec) DecodeKV(k, v []byte) (Entry, error) {
	panic("implement me")
}

func (u UniqueKeyCodec) EncodeKV(entry Entry) (k, v []byte, err error) {
	panic("implement me")
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
