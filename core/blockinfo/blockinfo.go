package blockinfo

import "google.golang.org/protobuf/types/known/timestamppb"

// BlockInfo represents basic block info independent of any specific Tendermint
// core version.
type BlockInfo interface {
	ChainID() string
	Height() int64
	Time() *timestamppb.Timestamp
	Hash() []byte
}
