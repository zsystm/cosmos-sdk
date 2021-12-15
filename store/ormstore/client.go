package ormstore

import (
	"github.com/cosmos/cosmos-sdk/orm/model/kvstore"
	"github.com/cosmos/cosmos-sdk/orm/model/ormhooks"
	"github.com/cosmos/cosmos-sdk/store/types"
)

type kvStoreBackend struct {
	store types.KVStore
}

func (k kvStoreBackend) Set(key, value []byte) error {
	k.store.Set(key, value)
	return nil
}

func (k kvStoreBackend) Delete(key []byte) error {
	k.store.Delete(key)
	return nil
}

func (k kvStoreBackend) CommitmentStore() kvstore.Store {
	return k
}

func (k kvStoreBackend) IndexStore() kvstore.Store {
	return k
}

func (k kvStoreBackend) ORMHooks() ormhooks.Hooks {
	//TODO implement me
	panic("implement me")
}

func (k kvStoreBackend) CommitmentStoreReader() kvstore.Reader {
	return k
}

func (k kvStoreBackend) IndexStoreReader() kvstore.Reader {
	return k
}

func (k kvStoreBackend) Get(key []byte) ([]byte, error) {
	x := k.store.Get(key)
	return x, nil
}

func (k kvStoreBackend) Has(key []byte) (bool, error) {
	x := k.store.Has(key)
	return x, nil
}

func (k kvStoreBackend) Iterator(start, end []byte) (kvstore.Iterator, error) {
	x := k.store.Iterator(start, end)
	return x, nil
}

func (k kvStoreBackend) ReverseIterator(start, end []byte) (kvstore.Iterator, error) {
	x := k.store.ReverseIterator(start, end)
	return x, nil
}

var _ kvstore.Backend = &kvStoreBackend{}
