package kv

import sdkstore "github.com/cosmos/cosmos-sdk/store"

type ReadStore interface {
	Get(key []byte) ([]byte, error)
	Has(key []byte) (bool, error)
	Iterator(start, end []byte) Iterator
	ReverseIterator(start, end []byte) Iterator
}

type Store interface {
	ReadStore
	Set(key, value []byte) error
	Delete(key []byte) error
}

type Iterator = sdkstore.Iterator
