package ormerrors

import "github.com/cosmos/cosmos-sdk/types/errors"

const codespace = "orm"

var (
	UnsupportedOperation = errors.New(codespace, 1, "unsupported operation")
)
