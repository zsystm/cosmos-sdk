package authn

import (
	"context"

	authnv1 "cosmossdk.io/api/cosmos/authn/v1"
	"cosmossdk.io/core/appmodule"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
)

type Keeper struct {
	authnv1.UnimplementedMsgServer
	authnv1.UnimplementedInternalServer

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
	acct.Credential = msg.NewCredential

	err = s.stateStore.AccountTable().Update(ctx, acct)
	if err != nil {
		return nil, err
	}

	err = s.appService.GetEventManager(ctx).Emit(&authnv1.EventSetCredential{
		Address:       addressStr,
		OldCredential: oldCredential,
		NewCredential: msg.NewCredential,
	})

	return &authnv1.MsgSetCredentialResponse{}, err
}

func (s Keeper) CreateAccount(ctx context.Context, request *authnv1.CreateAccountRequest) (*authnv1.CreateAccountResponse, error) {
	acc := &authnv1.Account{
		Address:    request.Address,
		Credential: request.Credential,
	}

	id, err := s.stateStore.AccountTable().InsertReturningId(ctx, acc)
	if err != nil {
		return nil, err
	}

	err = s.stateStore.AccountSequenceTable().Insert(ctx, &authnv1.AccountSequence{
		Id:  id,
		Seq: 0,
	})
	if err != nil {
		return nil, err
	}

	return &authnv1.CreateAccountResponse{AccountId: id}, nil
}

func (s Keeper) IncrementSeq(ctx context.Context, request *authnv1.IncrementSeqRequest) (*authnv1.IncrementSeqResponse, error) {
	accSeq, err := s.stateStore.AccountSequenceTable().Get(ctx, request.AccountId)
	if err != nil {
		return nil, err
	}

	accSeq.Seq = accSeq.Seq + 1
	err = s.stateStore.AccountSequenceTable().Update(ctx, accSeq)
	if err != nil {
		return nil, err
	}

	return &authnv1.IncrementSeqResponse{NewSeq: accSeq.Seq}, nil
}

var _ authnv1.MsgServer = &Keeper{}
var _ authnv1.InternalServer = &Keeper{}
