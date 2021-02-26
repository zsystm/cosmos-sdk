package ecdsa_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ecdsa"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/tendermint/tendermint/crypto"
)

func BenchmarkSecp256k1(b *testing.B) {
	skN, err := ecdsa.GenSecp256k1()
	if err != nil {
		b.Error(err)
	}
	pkN := skN.PubKey()
	skO := secp256k1.GenPrivKey()
	pkO := skO.PubKey()
	msg := crypto.CRandBytes(1000)
	b.ReportAllocs()

	b.Run("NEW", func(b *testing.B) {
		benchmarkSig(b, msg, skN, pkN)
	})

	b.Run("OLD", func(b *testing.B) {
		benchmarkSig(b, msg, skO, pkO)
	})
}

func benchmarkSig(b *testing.B, msg []byte, sk cryptotypes.PrivKey, pk cryptotypes.PubKey) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sig, err := sk.Sign(msg)
		if err != nil {
			b.Error(err)
		}
		// pk.VerifySignature(msg, sig)
		if !pk.VerifySignature(msg, sig) {
			b.Error("Verification failed")
		}
	}

}
