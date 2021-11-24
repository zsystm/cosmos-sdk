package ormkv

import "encoding/binary"

func AppendVarUint32(prefix []byte, id uint32) []byte {
	nsPrefixLen := len(prefix)
	res := make([]byte, nsPrefixLen+binary.MaxVarintLen32)
	copy(res, prefix)
	n := binary.PutUvarint(res[nsPrefixLen:], uint64(id))
	return res[:nsPrefixLen+n]
}
