package ormstore

import (
	"github.com/cosmos/cosmos-sdk/orm/model/kvstore"
	"github.com/cosmos/cosmos-sdk/store/types"
)

type kvStoreStore struct {
	store types.KVStore
}

type kvStoreWriter struct {
	*kvStoreStore
	batch []*writeEntry
}

func (k kvStoreWriter) Set(key, value []byte) error {
	k.batch = append(k.batch, &writeEntry{
		key:   key,
		value: value,
	})
	return nil
}

func (k kvStoreWriter) Delete(key []byte) error {
	k.batch = append(k.batch, &writeEntry{
		key:    key,
		delete: true,
	})
	return nil
}

func (k kvStoreWriter) CommitmentStoreWriter() kvstore.Writer {
	return k
}

func (k kvStoreWriter) IndexStoreWriter() kvstore.Writer {
	return k
}

func (k *kvStoreWriter) Write() error {
	for _, entry := range k.batch {
		if entry.delete {
			k.store.Delete(entry.key)
		} else {
			k.store.Set(entry.key, entry.value)
		}
	}
	k.batch = nil
	return nil
}

func (k kvStoreWriter) Close() {
	k.batch = nil
}

type writeEntry struct {
	key    []byte
	value  []byte
	delete bool
}

func (k kvStoreStore) CommitmentStoreReader() kvstore.Reader {
	return k
}

func (k kvStoreStore) IndexStoreReader() kvstore.Reader {
	return k
}

func (k *kvStoreStore) NewWriter() kvstore.IndexCommitmentStoreWriter {
	entries := make([]*writeEntry, 0, 16) // default capacity of 16 for index key updates
	return &kvStoreWriter{
		kvStoreStore: k,
		batch:        entries,
	}
}

func (k kvStoreStore) Get(key []byte) ([]byte, error) {
	x := k.store.Get(key)
	return x, nil
}

func (k kvStoreStore) Has(key []byte) (bool, error) {
	x := k.store.Has(key)
	return x, nil
}

func (k kvStoreStore) Iterator(start, end []byte) (kvstore.Iterator, error) {
	x := k.store.Iterator(start, end)
	return x, nil
}

func (k kvStoreStore) ReverseIterator(start, end []byte) (kvstore.Iterator, error) {
	x := k.store.ReverseIterator(start, end)
	return x, nil
}

var _ kvstore.Reader = &kvStoreStore{}
var _ kvstore.IndexCommitmentStore = &kvStoreStore{}
