package coin

import "github.com/cosmos/cosmos-sdk/base/math"

type DecCoin struct {
	Denom  string
	Amount math.Dec
}

type DecCoins map[string]math.Dec
