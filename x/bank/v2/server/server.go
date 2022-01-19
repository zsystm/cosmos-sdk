package server

import (
	"fmt"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"

	bankv1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/bank/v1beta1"
	bankv2alpha1 "github.com/cosmos/cosmos-sdk/api/cosmos/bank/v2alpha1"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/types/address"
)

// ModuleSchemaDescriptor declares a module's db statically
var ModuleSchemaDescriptor = ormdb.ModuleSchema{
	FileDescriptors: map[uint32]protoreflect.FileDescriptor{
		1: bankv2alpha1.File_cosmos_bank_v2alpha1_state_proto,
	},
}

type server struct {
	bankv1beta1.UnimplementedMsgServer
	bankv1beta1.UnimplementedQueryServer

	addressCodec address.Codec

	db                       ormdb.DB
	balanceTable             ormtable.Table
	balanceAddressDenomIndex ormtable.UniqueIndex
	balanceDenomAddressIndex ormtable.Index

	supplyTable      ormtable.Table
	supplyDenomIndex ormtable.UniqueIndex
}

// NewServer creates a new server.
//
// db is derived from a store key, codec, and ModuleSchemaDescriptor. It
// would be provided at the framework level using a one-per-scope providers (see the container module).
func NewServer(db ormdb.DB, codec address.Codec) (*server, error) {
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
