package keyring

import (
	"github.com/99designs/keyring"

	"cosmossdk.io/crypto/v2"
)

//type Keyring interface {
//	FindByName(name string) crypto.Credential
//	FindByAddressString(addressString string) crypto.Credential
//	FindByAddressBytes(addressBytes []byte) crypto.Credential
//	FindByPubKey(proto.Message) crypto.Credential
//}

type Keyring struct {
	kr keyring.Keyring
}

func (k *Keyring) KeyByName(name string) crypto.Credential {
	//k.kr.Get(name)
	panic("TODO")
}
