package address

import "github.com/cosmos/cosmos-sdk/base/state"

type Address []byte

func ParseAddress(ctx state.Context, str string) (Address, error) {
	panic("TODO")
}

func (addr Address) String(ctx state.Context) (string, error) {
	panic("TODO")
}
