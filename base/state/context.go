//Package state contains core types for state-machine logic
package state

import (
	"context"
	"time"

	"github.com/cosmos/cosmos-sdk/store/types"
)

type Context interface {
	context.Context
	BlockHeight() int64
	BlockTime() time.Time
	ChainID() string
	TxBytes() []byte
	KVStore(key types.StoreKey) types.KVStore
}
