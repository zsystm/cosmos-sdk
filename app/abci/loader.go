package abci

import (
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

type appLoader struct {
	abcitypes.Application
	configFileName string
}

func (a appLoader) Commit() abcitypes.ResponseCommit {
	//TODO implement me
	panic("implement me")
}

var _ abcitypes.Application = &appLoader{}
