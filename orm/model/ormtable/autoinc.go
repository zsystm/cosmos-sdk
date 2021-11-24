package ormtable

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/orm/backend/kv"
	"github.com/cosmos/cosmos-sdk/orm/encoding/ormkv"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
)

type AutoIncrementTable struct {
	*TableImpl
	autoIncField protoreflect.FieldDescriptor
	seqCodec     *ormkv.SeqCodec
}

func (s *AutoIncrementTable) Save(store kv.IndexCommitmentStore, message proto.Message, mode SaveMode) error {
	messageRef := message.ProtoReflect()
	val := messageRef.Get(s.autoIncField).Uint()
	if val == 0 {
		if mode == SAVE_MODE_UPDATE {
			return ormerrors.PrimaryKeyInvalidOnUpdate
		}

		mode = SAVE_MODE_CREATE
		key, err := s.nextSeqValue(store.IndexStore())
		if err != nil {
			return err
		}

		messageRef.Set(s.autoIncField, protoreflect.ValueOfUint64(key))
	} else {
		if mode == SAVE_MODE_CREATE {
			return ormerrors.AutoIncrementKeyAlreadySet
		}

		mode = SAVE_MODE_UPDATE
	}
	return s.TableImpl.Save(store, message, mode)
}

func (s *AutoIncrementTable) nextSeqValue(kv kv.Store) (uint64, error) {
	bz, err := kv.Get(s.seqCodec.Prefix)
	if err != nil {
		return 0, err
	}

	seq, err := s.seqCodec.DecodeValue(bz)
	if err != nil {
		return 0, err
	}

	seq++
	err = kv.Set(s.seqCodec.Prefix, s.seqCodec.EncodeValue(seq))
	return seq, err
}
