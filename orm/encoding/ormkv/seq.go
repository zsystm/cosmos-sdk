package ormkv

import (
	"bytes"
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type SeqCodec struct {
	TableName protoreflect.FullName
	Prefix    []byte
}

func (s SeqCodec) EncodeValue(seq uint64) (v []byte) {
	bz := make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(bz, seq)
	return bz[:n]
}

func (s SeqCodec) DecodeKV(k, v []byte) (Entry, error) {
	if !bytes.Equal(k, s.Prefix) {
		return nil, ormerrors.UnexpectedDecodePrefix
	}

	x, err := s.DecodeValue(v)
	if err != nil {
		return nil, err
	}

	return SeqEntry{
		TableName: s.TableName,
		Value:     x,
	}, nil
}

func (s SeqCodec) EncodeKV(entry Entry) (k, v []byte, err error) {
	seqEntry, ok := entry.(SeqEntry)
	if !ok {
		return nil, nil, ormerrors.BadDecodeEntry
	}

	if seqEntry.TableName != s.TableName {
		return nil, nil, ormerrors.BadDecodeEntry
	}

	return s.Prefix, s.EncodeValue(seqEntry.Value), nil
}

func (s SeqCodec) DecodeValue(v []byte) (uint64, error) {
	if len(v) == 0 {
		return 0, nil
	}
	return binary.ReadUvarint(bytes.NewReader(v))
}
