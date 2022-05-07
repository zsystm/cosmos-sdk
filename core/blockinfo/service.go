package blockinfo

import (
	"context"
)

// Service is a type which retrieves basic block info from a context independent
// of any specific Tendermint core version. Modules which need a specific
// Tendermint header should use a different service and should expect to need
// to update whenever Tendermint makes any changes.
type Service interface {
	GetBlockInfo(ctx context.Context) BlockInfo
}
