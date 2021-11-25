package ormkv

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
)

type SingletonKey struct {
	Prefix  []byte
	MsgType protoreflect.MessageType
}

func (s SingletonKey) DecodeKV(_, v []byte) (Entry, error) {
	msg := s.MsgType.New().Interface()
	err := proto.Unmarshal(v, msg)
	return PrimaryKeyEntry{Value: msg}, err
}

func (s SingletonKey) EncodeKV(entry Entry) (k, v []byte, err error) {
	pEntry, ok := entry.(PrimaryKeyEntry)
	if !ok {
		return nil, nil, ormerrors.BadDecodeEntry
	}

	if len(pEntry.Key) != 0 {
		return nil, nil, ormerrors.BadDecodeEntry.Wrap("singleton entry shouldn't have non-empty a key")
	}

	bz, err := proto.Marshal(pEntry.Value)
	if err != nil {
		return nil, nil, err
	}

	return s.Prefix, bz, nil
}
