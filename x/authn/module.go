package authn

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/address"

	modulev1 "cosmossdk.io/api/cosmos/authn/module/v1"
	authnv1 "cosmossdk.io/api/cosmos/authn/v1"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"

	"cosmossdk.io/core/appmodule"
)

func init() {
	appmodule.Register(&modulev1.Module{},
		appmodule.Provide(ProvideApp, ProvideAddressCodec),
	)
}

func ProvideApp(
	config *modulev1.Module,
	service appmodule.Service, db ormdb.ModuleDB) (*appmodule.Handler, error) {
	keeper, err := NewKeeper(config.Bech32Prefix, service, db)
	if err != nil {
		return nil, err
	}

	handler := &appmodule.Handler{}
	db.RegisterGenesisHandlers(handler)
	authnv1.RegisterMsgServer(handler, keeper)
	authnv1.RegisterInternalServer(handler, keeper)

	return handler, nil
}

func ProvideAddressCodec(config *modulev1.Module) (address.Codec, error) {
	if config.Bech32Prefix == "" {
		return nil, fmt.Errorf("missing bech32_prefix")
	}

	return NewBech32Codec(config.Bech32Prefix), nil
}
