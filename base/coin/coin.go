package coin

import "github.com/cosmos/cosmos-sdk/base/math"

type Coin struct {
	Denom  string
	Amount math.Int
}

type Coins map[string]math.Int
