package runtime

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
)

type App struct {
	*baseapp.BaseApp
}

func (a App) RegisterAPIRoutes(server *api.Server, config config.APIConfig) {}

func (a App) RegisterTxService(clientCtx client.Context) {}

func (a App) RegisterTendermintService(clientCtx client.Context) {}

var _ servertypes.Application = &App{}
