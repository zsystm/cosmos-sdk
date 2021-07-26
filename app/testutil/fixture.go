package testutil

import (
	"time"

	"github.com/cosmos/cosmos-sdk/app"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/container"
	"github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

type UnitTextFixture interface {
	InitGenesis([]byte)
	WaitForNextBlock() int64
	DeliverTx([]byte)
}

type unitTestFixture struct {
	baseapp *baseapp.BaseApp
}

func (u unitTestFixture) InitGenesis(genesis []byte) {
	_ = u.baseapp.InitChain(types.RequestInitChain{
		Time:            time.Time{},
		ChainId:         "",
		ConsensusParams: nil,
		Validators:      nil,
		AppStateBytes:   genesis,
		InitialHeight:   0,
	})
	_ = u.baseapp.BeginBlock(types.RequestBeginBlock{
		Hash:                nil,
		Header:              tmproto.Header{Height: u.baseapp.LastBlockHeight() + 1},
		LastCommitInfo:      types.LastCommitInfo{},
		ByzantineValidators: nil,
	})
}

func (u unitTestFixture) WaitForNextBlock() int64 {
	u.baseapp.EndBlock(types.RequestEndBlock{Height: u.baseapp.LastBlockHeight() + 1})
	u.baseapp.Commit()
	h := u.baseapp.LastBlockHeight() + 1
	u.baseapp.BeginBlock(types.RequestBeginBlock{
		Hash:                nil,
		Header:              tmproto.Header{Height: h},
		LastCommitInfo:      types.LastCommitInfo{},
		ByzantineValidators: nil,
	})
	return h
}

func (u unitTestFixture) DeliverTx(bz []byte) {
	u.baseapp.DeliverTx(types.RequestDeliverTx{Tx: bz})
}

var _ UnitTextFixture = &unitTestFixture{}

var ProvideUnitTestFixture = container.Options(
	app.ProvideApp,
	container.Provide(func(baseApp *baseapp.BaseApp) UnitTextFixture {
		return &unitTestFixture{baseapp: baseApp}
	}))
