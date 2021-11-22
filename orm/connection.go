package orm

import "context"

type ReadStoreConnection interface {
	OpenRead(context.Context) (ReadStore, error)
}

type StoreConnection interface {
	ReadStoreConnection
	Open(context.Context) (Store, error)
}
