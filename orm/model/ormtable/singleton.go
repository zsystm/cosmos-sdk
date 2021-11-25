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

	typeResolver    TypeResolver
	customValidator func(proto.Message) error
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

func (s *Singleton) decodeJson(reader io.Reader) (proto.Message, error) {
	bz, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	msg := s.MsgType.New().Interface()
	err = protojson.UnmarshalOptions{Resolver: s.typeResolver}.Unmarshal(bz, msg)
	return msg, err
}

func (s *Singleton) ValidateJSON(reader io.Reader) error {
	msg, err := s.decodeJson(reader)
	if err != nil {
		return err
	}

	if s.customValidator != nil {
		return s.customValidator(msg)
	} else {
		return DefaultValidator(msg)
	}
}

func (s *Singleton) ImportJSON(store kv.IndexCommitmentStore, reader io.Reader) error {
	msg, err := s.decodeJson(reader)
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

var _ Table = &Singleton{}
