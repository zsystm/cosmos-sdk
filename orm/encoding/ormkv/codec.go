package ormkv

type Codec interface {
	DecodeKV(k, v []byte) (Entry, error)
	EncodeKV(entry Entry) (k, v []byte, err error)
}
