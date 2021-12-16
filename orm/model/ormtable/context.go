package ormtable

import (
	"github.com/cosmos/cosmos-sdk/orm/model/kvstore"
)

// ReadContext defines the type used for read-only ORM operations.
type ReadContext interface {
	// CommitmentStoreReader returns the reader for the commitment store.
	CommitmentStoreReader() kvstore.Reader

	// IndexStoreReader returns the reader for the index store.
	IndexStoreReader() kvstore.Reader
}

// Context defines the type used for read-write ORM operations.
// Unlike ReadContext, write access to the underlying kv-store
// is hidden so that this can be fully encapsulated by the ORM.
type Context interface {
	ReadContext
	getCommitmentStore() kvstore.Store
	getIndexStore() kvstore.Store
	getHooks() Hooks
}

// ReadContextOptions defines options for creating a ReadContext.
// Read context can optionally define two stores - a commitment store
// that is backed by a merkle tree and an index store that isn't.
// If the index store is not defined, the commitment store will be
// used for all operations.
type ReadContextOptions struct {

	// CommitmentStoreReader is a reader for the commitment store.
	CommitmentStoreReader kvstore.Reader

	// IndexStoreReader is an optional reader for the index store.
	// If it is nil the CommitmentStoreReader will be used.
	IndexStoreReader kvstore.Reader
}

type readContext struct {
	commitmentReader kvstore.Reader
	indexReader      kvstore.Reader
}

func (r readContext) CommitmentStoreReader() kvstore.Reader {
	return r.commitmentReader
}

func (r readContext) IndexStoreReader() kvstore.Reader {
	return r.indexReader
}

// NewReadContext creates a new ReadContext.
func NewReadContext(options ReadContextOptions) ReadContext {
	indexReader := options.IndexStoreReader
	if indexReader == nil {
		indexReader = options.CommitmentStoreReader
	}
	return &readContext{
		commitmentReader: options.CommitmentStoreReader,
		indexReader:      indexReader,
	}
}

type context struct {
	commitmentStore kvstore.Store
	indexStore      kvstore.Store
	hooks           Hooks
}

func (c context) CommitmentStoreReader() kvstore.Reader {
	return c.commitmentStore
}

func (c context) IndexStoreReader() kvstore.Reader {
	return c.indexStore
}

func (c context) getCommitmentStore() kvstore.Store {
	return c.commitmentStore
}

func (c context) getIndexStore() kvstore.Store {
	return c.indexStore
}

func (c context) getHooks() Hooks {
	return c.hooks
}

// ContextOptions defines options for creating a Context.
// Context can optionally define two stores - a commitment store
// that is backed by a merkle tree and an index store that isn't.
// If the index store is not defined, the commitment store will be
// used for all operations.
type ContextOptions struct {

	// CommitmentStore is the commitment store.
	CommitmentStore kvstore.Store

	// IndexStore is the optional index store.
	// If it is nil the CommitmentStore will be used.
	IndexStore kvstore.Store

	// Hooks are optional hooks into ORM insert, update and delete operations.
	Hooks Hooks
}

// NewContext creates a new Context.
func NewContext(options ContextOptions) Context {
	indexStore := options.IndexStore
	if indexStore == nil {
		indexStore = options.CommitmentStore
	}
	return &context{
		commitmentStore: options.CommitmentStore,
		indexStore:      indexStore,
		hooks:           options.Hooks,
	}
}
