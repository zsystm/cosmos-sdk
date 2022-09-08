package keyring

import "github.com/gogo/protobuf/proto"

type Keyring interface {
	KeyByName(name string) SigningKey
	KeyByAddressString(addressString string) SigningKey
	KeyByAddressBytes(addressBytes []byte) SigningKey
	KeyByPubKey(proto.Message) SigningKey
}

type SigningKey interface {
	PublicKey() PublicKey
	Sign([]byte) ([]byte, error)
}

type PublicKey interface {
	proto.Message
	VerifySignature(msg []byte, sig []byte) bool
	Address() []byte
}
