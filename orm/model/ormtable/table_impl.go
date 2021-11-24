package ormtable

import (
	"encoding/json"
	io "io"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/orm/backend/kv"
	"github.com/cosmos/cosmos-sdk/orm/encoding/ormkv"
	"github.com/cosmos/cosmos-sdk/orm/model/ormindex"
)

type TableImpl struct {
	*ormindex.PrimaryKey
	indexers              []ormindex.Indexer
	indexes               []ormindex.Index
	indexesByFields       map[Fields]ormindex.Index
	uniqueIndexesByFields map[Fields]ormindex.UniqueIndex
	indexesById           map[uint32]ormindex.Index
	prefix                []byte
}

func (t TableImpl) Save(store kv.IndexCommitmentStore, message proto.Message, mode SaveMode) error {
	mref := message.ProtoReflect()
	pkValues, pk, err := t.EncodeFromMessage(mref)
	if err != nil {
		return err
	}

	existing := mref.New().Interface()
	haveExisting, err := t.GetByKeyBytes(store, pk, pkValues, existing)
	if err != nil {
		return err
	}

	if haveExisting {
		if mode == SAVE_MODE_CREATE {
			return ormerrors.PrimaryKeyConstraintViolation.Wrapf("%q", mref.Descriptor().FullName())
		}
	} else {
		if mode == SAVE_MODE_UPDATE {
			return ormerrors.NotFoundOnUpdate.Wrapf("%q", mref.Descriptor().FullName())
		}
	}

	// temporarily clear primary key
	t.ClearValues(mref)

	// store object
	bz, err := proto.MarshalOptions{Deterministic: true}.Marshal(message)
	err = store.CommitmentStore().Set(pk, bz)
	if err != nil {
		return err
	}

	// set primary key again

	t.SetValues(mref, pkValues)

	// set indexes
	indexStore := store.IndexStore()
	if !haveExisting {
		for _, idx := range t.indexers {
			err = idx.OnCreate(indexStore, mref)
			if err != nil {
				return err
			}

		}
	} else {
		existingMref := existing.ProtoReflect()
		for _, idx := range t.indexers {
			err = idx.OnUpdate(indexStore, mref, existingMref)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (t TableImpl) Delete(store kv.IndexCommitmentStore, primaryKey []protoreflect.Value) error {
	pk, err := t.Encode(primaryKey)
	if err != nil {
		return err
	}

	msg := t.Type.New().Interface()
	found, err := t.GetByKeyBytes(store, pk, primaryKey, msg)
	if err != nil {
		return err
	}

	if !found {
		return nil
	}

	// delete object
	err = store.CommitmentStore().Delete(pk)
	if err != nil {
		return err
	}

	// clear indexes
	mref := msg.ProtoReflect()
	indexStore := store.IndexStore()
	for _, idx := range t.indexers {
		err := idx.OnDelete(indexStore, mref)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t TableImpl) GetIndex(fields Fields) ormindex.Index {
	return t.indexesByFields[fields]
}

func (t TableImpl) GetUniqueIndex(fields Fields) ormindex.UniqueIndex {
	return t.uniqueIndexesByFields[fields]
}

func (t TableImpl) Indexes() []ormindex.Index {
	return t.indexes
}

func (t TableImpl) Decode(k []byte, v []byte) (ormkv.Entry, error) {
	panic("implement me")
}

func (t TableImpl) DefaultJSON() json.RawMessage {
	return json.RawMessage("[]")
}

func (t TableImpl) ValidateJSON(reader io.Reader) error {
	panic("implement me")
}

func (t TableImpl) ImportJSON(store kv.IndexCommitmentStore, reader io.Reader) error {
	panic("implement me")
}

func (t TableImpl) ExportJSON(store kv.IndexCommitmentReadStore, writer io.Writer) error {
	panic("implement me")
}

var _ Table = &TableImpl{}
