package authn

import (
	"context"

	"github.com/cosmos/cosmos-sdk/types/address"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	modulev1 "cosmossdk.io/api/cosmos/authn/module/v1"
	authnv1 "cosmossdk.io/api/cosmos/authn/v1"

	"cosmossdk.io/core/appmodule"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
)

type Keeper struct {
	authnv1.UnimplementedMsgServer
	authnv1.UnimplementedInternalServer
	authnv1.UnimplementedAdminServer

	addressCodec Bech32Codec
	adminModules map[string]bool
	appService   appmodule.Service
	stateStore   authnv1.StateStore
}

func NewKeeper(config *modulev1.Module, appService appmodule.Service, db ormdb.ModuleDB) (*Keeper, error) {
	stateStore, err := authnv1.NewStateStore(db)
	if err != nil {
		return nil, err
	}

	adminModules := map[string]bool{}
	for _, module := range config.AdminModules {
		adminModules[module] = true
	}

	return &Keeper{
		addressCodec: NewBech32Codec(config.Bech32Prefix),
		appService:   appService,
		stateStore:   stateStore,
		adminModules: adminModules,
	}, nil
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
	callingModule := s.appService.InternalServiceCaller(ctx)
	if !s.adminModules[callingModule] {
		return nil, status.Errorf(codes.PermissionDenied, "%s is not an admin module", callingModule)
	}

	acc := &authnv1.Account{
		Address:    request.Address,
		Credential: request.Credential,
	}

	id, err := s.stateStore.AccountTable().InsertReturningId(ctx, acc)
	if err != nil {
		return nil, err
	}

	err = s.stateStore.AccountSequenceTable().Insert(ctx, &authnv1.AccountSequence{
		Address: request.Address,
		Seq:     0,
	})
	if err != nil {
		return nil, err
	}

	return &authnv1.CreateAccountResponse{AccountId: id}, nil
}

func (s Keeper) CreateModuleAccount(ctx context.Context, request *authnv1.CreateModuleAccountRequest) (*authnv1.CreateModuleAccountResponse, error) {
	callingModule := s.appService.InternalServiceCaller(ctx)
	addrBytes := address.Module(callingModule, request.DerivationPath)
	acc := &authnv1.Account{
		Address:    addrBytes,
		Credential: request.Credential,
	}

	id, err := s.stateStore.AccountTable().InsertReturningId(ctx, acc)
	if err != nil {
		return nil, err
	}

	err = s.stateStore.AccountSequenceTable().Insert(ctx, &authnv1.AccountSequence{
		Address: addrBytes,
		Seq:     0,
	})
	if err != nil {
		return nil, err
	}

	return &authnv1.CreateModuleAccountResponse{AccountId: id, Address: addrBytes}, nil
}

func (s Keeper) IncrementSeq(ctx context.Context, request *authnv1.IncrementSeqRequest) (*authnv1.IncrementSeqResponse, error) {
	callingModule := s.appService.InternalServiceCaller(ctx)
	if !s.adminModules[callingModule] {
		return nil, status.Errorf(codes.PermissionDenied, "%s is not an admin module", callingModule)
	}

	accSeq, err := s.stateStore.AccountSequenceTable().Get(ctx, request.Address)
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
var _ authnv1.AdminServer = &Keeper{}
