package authn

import (
	"context"

	"cosmossdk.io/core/appmodule"
	authnv1 "cosmossdk.io/x/authn/internal/cosmos/authn/v1"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
)

type Keeper struct {
	authnv1.UnimplementedMsgServer

	addressCodec Bech32Codec
	appService   appmodule.Service
	stateStore   authnv1.StateStore
}

func NewKeeper(bech32Prefix string, appService appmodule.Service, db ormdb.ModuleDB) (*Keeper, error) {
	stateStore, err := authnv1.NewStateStore(db)
	if err != nil {
		return nil, err
	}

	return &Keeper{addressCodec: NewBech32Codec(bech32Prefix), appService: appService, stateStore: stateStore}, nil
}

func (s Keeper) SetCredential(ctx context.Context, msg *authnv1.MsgSetCredential) (*authnv1.MsgSetCredentialResponse, error) {
	addressStr := msg.Address
	addressBz, err := s.addressCodec.StringToBytes(addressStr)
	if err != nil {
		return nil, err
	}

	acct, err := s.stateStore.AccountTable().GetByAddress(ctx, addressBz)
	if err != nil {
		return nil, err
	}

	oldCredential := acct.Credential
	newCredential := msg.NewCredential
	acct.Credential = newCredential

	err = s.stateStore.AccountTable().Update(ctx, acct)
	if err != nil {
		return nil, err
	}

	err = s.appService.GetEventManager(ctx).Emit(&authnv1.EventSetCredential{
		Address:       addressStr,
		OldCredential: oldCredential,
		NewCredential: newCredential,
	})

	return &authnv1.MsgSetCredentialResponse{}, err
}

var _ authnv1.MsgServer = &Keeper{}
