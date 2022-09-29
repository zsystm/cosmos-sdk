package crypto

import "google.golang.org/protobuf/proto"

type KeyProvider interface {
	CredentialFromProto(proto.Message) Credential
	PublicKeyFromProto(proto.Message) PublicKey
	PrivateKeyFromProto(proto.Message) PrivateKey
}

type Credential interface {
	AsProto() proto.Message
	Address() []byte
}

type PublicKey interface {
	Credential
	VerifySignature(msg []byte, sig []byte) bool
}

type PrivateKey interface {
	PublicKey() PublicKey
	AsProto() proto.Message
	Sign([]byte) ([]byte, error)
}
