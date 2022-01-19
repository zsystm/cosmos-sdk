package server

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"

	bankv1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/bank/v1beta1"
	bankv2alpha1 "github.com/cosmos/cosmos-sdk/api/cosmos/bank/v2alpha1"
	v1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	"github.com/cosmos/cosmos-sdk/orm/model/ormschema"
	"github.com/cosmos/cosmos-sdk/types/address"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ bankv1beta1.QueryServer = &server{}

func (s server) Balance(ctx context.Context, request *bankv1beta1.QueryBalanceRequest) (*bankv1beta1.QueryBalanceResponse, error) {
	addressBz, err := s.addressCodec.StringToBytes(request.Address)
	if err != nil {
		return nil, err
	}

	var balance bankv2alpha1.Balance
	found, err := s.balanceAddressDenomIndex.Get(ctx, &balance, addressBz, request.Denom)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &bankv1beta1.QueryBalanceResponse{Balance: &v1beta1.Coin{
		Denom:  request.Denom,
		Amount: balance.Amount,
	}}, nil
}

func (s server) AllBalances(ctx context.Context, request *bankv1beta1.QueryAllBalancesRequest) (*bankv1beta1.QueryAllBalancesResponse, error) {
	addressBz, err := s.addressCodec.StringToBytes(request.Address)
	if err != nil {
		return nil, err
	}

	res, err := ormtable.Paginate(
		s.balanceAddressDenomIndex,
		ctx,
		&ormtable.PaginationRequest{PageRequest: request.Pagination},
		ormlist.Prefix(addressBz),
	)
	if err != nil {
		return nil, err
	}

	return
}

func (s server) TotalSupply(ctx context.Context, request *bankv1beta1.QueryTotalSupplyRequest) (*bankv1beta1.QueryTotalSupplyResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s server) SupplyOf(ctx context.Context, request *bankv1beta1.QuerySupplyOfRequest) (*bankv1beta1.QuerySupplyOfResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s server) Params(ctx context.Context, request *bankv1beta1.QueryParamsRequest) (*bankv1beta1.QueryParamsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s server) DenomMetadata(ctx context.Context, request *bankv1beta1.QueryDenomMetadataRequest) (*bankv1beta1.QueryDenomMetadataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s server) DenomsMetadata(ctx context.Context, request *bankv1beta1.QueryDenomsMetadataRequest) (*bankv1beta1.QueryDenomsMetadataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s server) DenomOwners(ctx context.Context, request *bankv1beta1.QueryDenomOwnersRequest) (*bankv1beta1.QueryDenomOwnersResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s server) mustEmbedUnimplementedQueryServer() {
	//TODO implement me
	panic("implement me")
}

// NewServer creates a new server.
//
// db is derived from a store key, codec, and ModuleSchemaDescriptor. It
// would be provided at the framework level using a one-per-scope providers (see the container module).
func NewServer(db ormschema.DB, codec address.Codec) (*server, error) {
	s := &server{
		db:           db,
		addressCodec: codec,
	}

	var err error

	s.balanceTable, err = s.db.GetTable(&bankv2alpha1.Balance{})
	if err != nil {
		return nil, err
	}

	s.balanceAddressDenomIndex = s.balanceTable.GetUniqueIndex("address,denom")
	if s.balanceAddressDenomIndex == nil {
		return nil, fmt.Errorf("missing address,denom index")
	}

	s.balanceDenomAddressIndex = s.balanceTable.GetUniqueIndex("denom,address")
	if s.balanceDenomAddressIndex == nil {
		return nil, fmt.Errorf("missing denom,address index")
	}

	s.supplyTable, err = s.db.GetTable(&bankv2alpha1.Supply{})
	if err != nil {
		return nil, err
	}

	s.supplyDenomIndex = s.supplyTable.GetUniqueIndex("denom")
	if s.supplyDenomIndex == nil {
		return nil, fmt.Errorf("missing denom index")
	}

	return s, nil
}
