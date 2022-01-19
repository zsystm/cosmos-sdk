package server

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"

	bankv1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/bank/v1beta1"
	bankv2alpha1 "github.com/cosmos/cosmos-sdk/api/cosmos/bank/v2alpha1"
	"github.com/cosmos/cosmos-sdk/orm/model/ormschema"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/types/address"
)

// ModuleSchemaDescriptor declares a module's db statically
var ModuleSchemaDescriptor = ormschema.ModuleDescriptor{
	FileDescriptors: map[uint32]protoreflect.FileDescriptor{
		1: bankv2alpha1.File_cosmos_bank_v2alpha1_state_proto,
	},
}

type server struct {
	bankv1beta1.UnimplementedMsgServer
	bankv1beta1.UnimplementedQueryServer

	addressCodec address.Codec

	db                       ormschema.DB
	balanceTable             ormtable.Table
	balanceAddressDenomIndex ormtable.UniqueIndex
	balanceDenomAddressIndex ormtable.Index

	supplyTable      ormtable.Table
	supplyDenomIndex ormtable.UniqueIndex
}
