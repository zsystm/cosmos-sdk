package server

import (
	bankv1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/bank/v1beta1"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/store/ormstore"
	"github.com/cosmos/cosmos-sdk/types/address"
)

type server struct {
	bankv1beta1.UnimplementedMsgServer

	dbConnection ormstore.StoreKeyDB
	addressCodec address.Codec
}

var (
	addressDenomFields = ormtable.CommaSeparatedFieldNames("address,denom")
	denomAddressFields = ormtable.CommaSeparatedFieldNames("denom,address")
)
