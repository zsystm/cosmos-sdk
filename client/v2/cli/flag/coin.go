package flag

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/pflag"
	"google.golang.org/protobuf/reflect/protoreflect"

	basev1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

type coinType struct{}

func (c coinType) NewValue(context.Context, *Builder) pflag.Value {
	return &coinValue{}
}

func (c coinType) DefaultValue() string {
	return ""
}

type coinValue struct {
	coin *basev1beta1.Coin
}

func (c coinValue) Get() protoreflect.Value {
	return protoreflect.ValueOfMessage(c.coin.ProtoReflect())
}

func (c coinValue) String() string {
	//TODO implement me
	panic("implement me")
}

func (c *coinValue) Set(coinStr string) error {
	coinStr = strings.TrimSpace(coinStr)

	reDecCoin := regexp.MustCompile(fmt.Sprintf(`^(%s)%s(%s)$`, reDecAmt, reSpc, coinDenomRegex()))

	matches := reDecCoin.FindStringSubmatch(coinStr)
	if matches == nil {
		return DecCoin{}, fmt.Errorf("invalid decimal coin expression: %s", coinStr)
	}

	amountStr, denomStr := matches[1], matches[2]

	amount, err := NewDecFromStr(amountStr)
	if err != nil {
		return DecCoin{}, errors.Wrap(err, fmt.Sprintf("failed to parse decimal coin amount: %s", amountStr))
	}

	if err := ValidateDenom(denomStr); err != nil {
		return DecCoin{}, fmt.Errorf("invalid denom cannot contain spaces: %s", err)
	}

	return NewDecCoinFromDec(denomStr, amount), nil
}

func (c coinValue) Type() string {
	return "coin"
}
