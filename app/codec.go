package app

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/container"
)

type codecInputs struct {
	container.StructArgs

	CodecClosures []codecClosure `group:"codec"`
}

type codecClosure func(codectypes.TypeRegistry)

var CodecProvider = container.Provide(func(inputs codecInputs) (
	codectypes.TypeRegistry,
	codec.Codec,
	codec.ProtoCodecMarshaler,
	codec.BinaryCodec,
	codec.JSONCodec,
	*codec.LegacyAmino,
) {

	typeRegistry := codectypes.NewInterfaceRegistry()
	for _, closure := range inputs.CodecClosures {
		closure(typeRegistry)
	}
	cdc := codec.NewProtoCodec(typeRegistry)
	amino := codec.NewLegacyAmino()
	return typeRegistry, cdc, cdc, cdc, cdc, amino
})
