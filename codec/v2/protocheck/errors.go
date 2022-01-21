package protocheck

import "github.com/cosmos/cosmos-sdk/errors"

var codespace = "codec/protocheck"

var (
	ErrUnknownField            = errors.New(codespace, 1, "unknown protobuf field")
	ErrInterfaceNotImplemented = errors.New(codespace, 2, "message does not implement expected interface")
)
