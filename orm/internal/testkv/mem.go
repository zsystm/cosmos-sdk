package testkv

import (
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
)

// NewSplitMemBackend returns a Backend instance
// which uses two separate memory stores to simulate behavior when there
// are really two separate backing stores.
func NewSplitMemBackend() ormtable.Context {
	return ormtable.NewContext(ormtable.ContextOptions{
		CommitmentStore: dbm.NewMemDB(),
		IndexStore:      dbm.NewMemDB(),
	})
}

// NewSharedMemBackend returns a Backend instance
// which uses a single backing memory store to simulate legacy scenarios
// where only a single KV-store is available to modules.
func NewSharedMemBackend() ormtable.Context {
	return ormtable.NewContext(ormtable.ContextOptions{
		CommitmentStore: dbm.NewMemDB(),
		// commit store is automatically used as the index store
	})
}
