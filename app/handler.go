package app

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"
	"google.golang.org/grpc"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/container"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Handler struct {
	// Genesis
	InitGenesis func(context.Context, InitGenesisRequest) (InitGenesisResponse, error)
	//DefaultGenesis  func(codec.JSONCodec) json.RawMessage
	//ValidateGenesis func(codec.JSONCodec, client.TxEncodingConfig, json.RawMessage) error
	ExportGenesis func(sdk.Context, codec.JSONCodec) json.RawMessage

	// ABCI
	BeginBlocker  func(context.Context, BeginBlockRequest) (BeginBlockResponse, error)
	EndBlocker    func(context.Context, abci.RequestEndBlock) []ValidatorUpdate
	MsgServices   []ServiceImpl
	QueryServices []ServiceImpl

	// CLI
	QueryCommand *cobra.Command
	TxCommand    *cobra.Command
}

type InitGenesisRequest struct {
	JSONCodec codec.JSONCodec
	Content   json.RawMessage
}

type InitGenesisResponse struct {
	ValidatorUpdates []ValidatorUpdate
}

type BeginBlockRequest struct {
	Hash []byte
	// TODO:
	// Header              types1.Header  `protobuf:"bytes,2,opt,name=header,proto3" json:"header"`
	// LastCommitInfo      LastCommitInfo `protobuf:"bytes,3,opt,name=last_commit_info,json=lastCommitInfo,proto3" json:"last_commit_info"`
	// ByzantineValidators []Evidence     `protobuf:"bytes,4,rep,name=byzantine_validators,json=byzantineValidators,proto3" json:"byzantine_validators"`
}

type BeginBlockResponse struct {
}

type ValidatorUpdate struct {
	PubKey cryptotypes.PubKey
	Power  int64
}

type ServiceImpl struct {
	Desc *grpc.ServiceDesc
	Impl interface{}
}

func (Handler) IsOnePerScopeType() {}

var _ container.OnePerScopeType = Handler{}
