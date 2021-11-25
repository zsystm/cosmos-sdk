package ormtable

import (
	"encoding/json"
	io "io"

	"github.com/cosmos/cosmos-sdk/orm/encoding/ormkv"

	"google.golang.org/protobuf/encoding/protojson"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/orm/backend/kv"
	"github.com/cosmos/cosmos-sdk/orm/model/ormindex"
)

type Singleton struct {
	*ormindex.SingletonIndex
}

func (s Singleton) Save(store kv.IndexCommitmentStore, message proto.Message, _ SaveMode) error {
	bz, err := proto.Marshal(message)
	if err != nil {
		return err
	}
	return store.CommitmentStore().Set(s.Prefix, bz)
}

func (s Singleton) Delete(store kv.IndexCommitmentStore, _ []protoreflect.Value) error {
	return store.CommitmentStore().Delete(s.Prefix)
}

func (s Singleton) GetIndex(fields ormkv.Fields) ormindex.Index {
	if fields.String() != "" {
		return nil
	}
	return s.SingletonIndex
}

func (s Singleton) GetUniqueIndex(fields ormkv.Fields) ormindex.UniqueIndex {
	if fields.String() != "" {
		return nil
	}
	return s.SingletonIndex
}

func (s Singleton) Indexes() []ormindex.Index {
	return []ormindex.Index{s}
}

func (s *Singleton) DefaultJSON() json.RawMessage {
	msg := s.MsgType.New().Interface()
	bz, err := protojson.MarshalOptions{}.Marshal(msg)
	if err != nil {
		return json.RawMessage("{}")
	}
	return bz
}

func (s *Singleton) ValidateJSON(reader io.Reader) error {
	panic("implement me")
}

func (s *Singleton) ImportJSON(store kv.IndexCommitmentStore, reader io.Reader) error {
	bz, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	msg := s.MsgType.New().Interface()
	err = protojson.Unmarshal(bz, msg)
	if err != nil {
		return err
	}

	return s.Save(store, msg, SAVE_MODE_DEFAULT)
}

func (s *Singleton) ExportJSON(store kv.IndexCommitmentReadStore, writer io.Writer) error {
	msg := s.MsgType.New().Interface()
	found, err := s.Get(store, nil, msg)
	if err != nil {
		return err
	}

	var bz []byte
	if !found {
		bz = s.DefaultJSON()
	} else {
		bz, err = protojson.Marshal(msg)
		if err != nil {
			return err
		}
	}

	_, err = writer.Write(bz)
	return err
}
