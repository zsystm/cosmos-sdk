// ECDSA package implements Cosmos-SDK compatible ECDSA public and private key. The keys
// can be protobuf serialized and packed in Any.
// Currently supported keys are:
// + secp256r1
package ecdsa

import (
	"crypto/elliptic"
	"fmt"
	"math/big"
	math_bits "math/bits"

	secp256k1pkg "github.com/fomichev/secp256k1"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

const (
	// PubKeySize is is the size, in bytes, of public keys as used in this package.
	PubKeySize = 32 + 1 + 1
	// PrivKeySize is the size, in bytes, of private keys as used in this package.
	PrivKeySize = 32 + 1
)

var secp256r1, secp256k1 elliptic.Curve
var curveNames map[elliptic.Curve]string
var curveTypes map[elliptic.Curve]byte
var curveTypesRev map[byte]elliptic.Curve

// Protobuf Bytes size - this computation is based on gogotypes.BytesValue.Sizee implementation
var sovPubKeySize = 1 + PubKeySize + sovKeys(PubKeySize)
var sovPrivKeySize = 1 + PrivKeySize + sovKeys(PrivKeySize)

func init() {
	secp256r1 = elliptic.P256()
	secp256k1 = secp256k1pkg.SECP256K1()
	// initSECP256K1() // there is something wrong with the stdlib curve operations

	// PubKeySize is ceil of field bit size + 1 for the sign + 1 for the type
	expected := (secp256r1.Params().BitSize+7)/8 + 2
	if expected != PubKeySize {
		panic(fmt.Sprintf("Wrong PubKeySize=%d, expecting=%d", PubKeySize, expected))
	}

	curveNames = map[elliptic.Curve]string{
		secp256r1: "secp256r1",
		secp256k1: "secp256k1",
	}
	curveTypes = map[elliptic.Curve]byte{
		// 0 reserved
		secp256r1: 1,
		secp256k1: 2,
	}
	curveTypesRev = map[byte]elliptic.Curve{}
	for c, b := range curveTypes {
		curveTypesRev[b] = c
	}
}

func initSECP256K1() {
	// http://www.secg.org/sec2-v2.pdf
	cp := &elliptic.CurveParams{Name: "secp256k1"}
	cp.P, _ = new(big.Int).SetString("fffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f", 16)
	cp.N, _ = new(big.Int).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
	cp.B, _ = new(big.Int).SetString("0000000000000000000000000000000000000000000000000000000000000007", 16)
	cp.Gx, _ = new(big.Int).SetString("79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798", 16)
	cp.Gy, _ = new(big.Int).SetString("483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8", 16)
	cp.BitSize = 256
	secp256k1 = cp
}

// RegisterInterfaces adds ecdsa PubKey to pubkey registry
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &ecdsaPK{})
}

func sovKeys(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
