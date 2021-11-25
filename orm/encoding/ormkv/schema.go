package ormkv

import (
	"bytes"

	"github.com/gogo/protobuf/proto"

	"google.golang.org/protobuf/types/descriptorpb"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
)

type SchemaKeyCodec struct {
	Prefix []byte
}

func (s SchemaKeyCodec) DecodeKV(k, v []byte) (Entry, error) {
	if !bytes.Equal(k, s.Prefix) {
		return nil, ormerrors.UnexpectedDecodePrefix
	}

	var fd descriptorpb.FileDescriptorProto
	err := proto.Unmarshal(v, &fd)
	if err != nil {
		return nil, ormerrors.BadDecodeEntry.Wrapf("bad FileDescriptorProto %v", err)
	}

	return SchemaEntry{
		FileDescriptor: nil,
	}, nil
}

func (s SchemaKeyCodec) EncodeKV(entry Entry) (k, v []byte, err error) {
	v, err = proto.Marshal(entry.(SchemaEntry).FileDescriptor)
	return s.Prefix, v, nil
}

var _ Codec = &SchemaKeyCodec{}
