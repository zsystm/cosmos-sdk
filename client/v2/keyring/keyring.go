package keyring

import "google.golang.org/protobuf/proto"

type Keyring interface {
	KeyByName(name string) SigningKey
	KeyByAddressString(addressString string) SigningKey
	KeyByAddressBytes(addressBytes []byte) SigningKey
	KeyByPubKey(proto.Message) SigningKey
}

type KeyProvider interface {
	PublicKeyFromProto(proto.Message) PublicKey
	SigningKeyFromProto(proto.Message) PrivateKey
}

type PublicKey interface {
	AsProto() proto.Message
	VerifySignature(msg []byte, sig []byte) bool
	Address() []byte
}

type SigningKey interface {
	PublicKey() PublicKey
	Sign([]byte) ([]byte, error)
}

type PrivateKey interface {
	SigningKey
	AsProto() proto.Message
}
