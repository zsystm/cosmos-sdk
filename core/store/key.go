package store

import "context"

type Type int

// Key represents a unique, non-forgeable handle to a KVStore.
type Key interface {
	KVStore(context.Context) KVStore
	TransientStore(context.Context) KVStore
	MemoryStore(context.Context) KVStore
}
