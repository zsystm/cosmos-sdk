package app

import (
	"github.com/cosmos/cosmos-sdk/container"
)

var ProvideApp = container.Options(
	BaseAppProvider,
	StoreKeyProvider,
	CodecProvider,
)
