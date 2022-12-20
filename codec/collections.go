package codec

import (
	"cosmossdk.io/collections"
	gogoproto "github.com/cosmos/gogoproto/proto"
	gogotypes "github.com/cosmos/gogoproto/types"
	"log"
	"strconv"
)

var ProtoBoolValue collections.ValueCodec[bool] = boolProto{}

type boolProto struct{}

func (boolProto) Encode(value bool) ([]byte, error) {
	log.Printf("%v", value)
	return (&gogotypes.BoolValue{Value: value}).Marshal()
}

func (boolProto) Decode(b []byte) (bool, error) {
	x := new(gogotypes.BoolValue)
	err := x.Unmarshal(b)
	return x.Value, err
}

func (boolProto) Stringify(value bool) string {
	return strconv.FormatBool(value)
}

func (boolProto) ValueType() string {
	return gogoproto.MessageName(new(gogotypes.BoolValue))
}

type ProtoValueCodec[T any, PT interface {
	*T
	gogoproto.Message
}] struct {
	_   T
	cdc BinaryCodec
}

func (p *ProtoValueCodec[T, PT]) Encode(v T) ([]byte, error) {
	return p.cdc.Marshal(PT(&v))
}

func (p *ProtoValueCodec[T, PT]) Decode(b []byte) (v T, err error) {
	x := PT(new(T))
	err = p.cdc.Unmarshal(b, x)
	if err != nil {
		return v, err
	}
	return *x, nil
}

func (p *ProtoValueCodec[T, PT]) Stringify(v T) string {
	return PT(&v).String()
}

func (p *ProtoValueCodec[T, PT]) ValueType() string { return gogoproto.MessageName(PT(new(T))) }

func NewProtoValueCodec[T any, PT interface {
	*T
	gogoproto.Message
}](cdc BinaryCodec) collections.ValueCodec[T] {
	return &ProtoValueCodec[T, PT]{
		cdc: cdc,
	}
}
