package context

import (
	"context"
	"github.com/cosmos/cosmos-sdk/app/event"
	"github.com/cosmos/cosmos-sdk/app/internal/gas"
	"time"
)

type Context struct {
	BaseCtx context.Context
	//ms            MultiStore
	//header        tmproto.Header
	ChainID    string
	TxBytes    []byte
	HeaderHash []byte
	//logger        log.Logger
	//voteInfo      []abci.VoteInfo
	GasMeter      gas.Meter
	BlockGasMeter gas.Meter
	//consParams    *tmproto.ConsensusParams
	EventManager event.Manager
}

func (c Context) Deadline() (deadline time.Time, ok bool) {
	return c.BaseCtx.Deadline()
}

func (c Context) Done() <-chan struct{} {
	return c.BaseCtx.Done()
}

func (c Context) Err() error {
	return c.BaseCtx.Err()
}

func (c Context) Value(key interface{}) interface{} {
	return c.BaseCtx.Value(key)
}

func (c Context) WithValue(key, value interface{}) context.Context {
	c.BaseCtx = context.WithValue(c.BaseCtx, key, value)
	return c
}

var _ context.Context = &Context{}
