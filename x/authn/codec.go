package authn

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Bech32Codec struct {
	bech32Prefix string
}

var _ address.Codec = &Bech32Codec{}

func NewBech32Codec(prefix string) Bech32Codec {
	return Bech32Codec{prefix}
}

// StringToBytes encodes text to bytes
func (bc Bech32Codec) StringToBytes(text string) ([]byte, error) {
	hrp, bz, err := bech32.DecodeAndConvert(text)
	if err != nil {
		return nil, err
	}

	if hrp != bc.bech32Prefix {
		return nil, status.Errorf(codes.InvalidArgument, "hrp does not match bech32Prefix")
	}

	if err := sdk.VerifyAddressFormat(bz); err != nil {
		return nil, err
	}

	return bz, nil
}

// BytesToString decodes bytes to text
func (bc Bech32Codec) BytesToString(bz []byte) (string, error) {
	text, err := bech32.ConvertAndEncode(bc.bech32Prefix, bz)
	if err != nil {
		return "", err
	}

	return text, nil
}
