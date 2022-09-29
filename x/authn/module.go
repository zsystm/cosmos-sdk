package authn

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/address"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"

	"cosmossdk.io/core/appmodule"
	authnv1 "cosmossdk.io/x/authn/internal/cosmos/authn/v1"
)

func (x *Module) ProvideApp(service appmodule.Service, db ormdb.ModuleDB) (*appmodule.Handler, error) {
	server, err := NewKeeper(x.Bech32Prefix, service, db)
	if err != nil {
		return nil, err
	}

	handler := &appmodule.Handler{}
	db.RegisterGenesisHandlers(handler)
	authnv1.RegisterMsgServer(handler, server)

	return handler, nil
}

func (x *Module) ProvideAddressCodec() (address.Codec, error) {
	if x.Bech32Prefix == "" {
		return nil, fmt.Errorf("missing bech32_prefix")
	}

	return NewBech32Codec(x.Bech32Prefix), nil
}
