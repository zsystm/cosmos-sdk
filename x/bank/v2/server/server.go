package server

import (
	bankv1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/bank/v1beta1"
	bankv2alpha1 "github.com/cosmos/cosmos-sdk/api/cosmos/bank/v2alpha1"
	"github.com/cosmos/cosmos-sdk/orm/model/ormschema"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

type server struct {
	bankv1beta1.UnimplementedMsgServer

	addressCodec address.Codec

	schema       *ormschema.ModuleSchema
	balanceTable ormtable.Table
	supplyTable  ormtable.Table
}

func NewServer(key *storetypes.KVStoreKey, codec address.Codec) (*server, error) {
	s := &server{
		addressCodec: codec,
	}

	var err error

	s.balanceTable, err = s.schema.GetTable(&bankv2alpha1.Balance{})
	if err != nil {
		return nil, err
	}

	s.supplyTable, err = s.schema.GetTable(&bankv2alpha1.Supply{})
	if err != nil {
		return nil, err
	}

	return s, nil
}
