package nft

import (
	"github.com/cosmos/cosmos-sdk/container"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	store "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/nft"
	nftkeeper "github.com/cosmos/cosmos-sdk/x/nft/keeper"
)

type Inputs struct {
	container.In

	StoreKey          *store.KVStoreKey
	Codec             codec.Codec
	AccountKeeper     authkeeper.AccountKeeper
	BankKeeper        nft.BankKeeper
	InterfaceRegistry codectypes.InterfaceRegistry
}

type Outputs struct {
	container.Out

	Keeper    nftkeeper.Keeper
	AppModule module.AppModuleWiringWrapper
}

func Provide(inputs Inputs) (Outputs, error) {
	k := nftkeeper.NewKeeper(inputs.StoreKey, inputs.Codec, inputs.AccountKeeper, inputs.BankKeeper)
	m := NewAppModule(inputs.Codec, k, inputs.AccountKeeper, inputs.BankKeeper, inputs.InterfaceRegistry)
	return Outputs{
		Keeper:    k,
		AppModule: module.AppModuleWiringWrapper{AppModule: m},
	}, nil
}
