package unittest

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/app"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/container"
)

type unitTestOutputs struct {
	SDKContextGetter func() sdk.Context
	ContextGetter    func() context.Context
}

var FixtureProvider = container.Options(
	app.ProvideApp,
	container.Provide(func(baseApp *baseapp.BaseApp) unitTestOutputs {
		return unitTestOutputs{
			SDKContextGetter: nil,
			ContextGetter:    nil,
		}
	}))
