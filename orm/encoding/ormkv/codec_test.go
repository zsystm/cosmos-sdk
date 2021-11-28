package ormkv_test

import (
	"bytes"
	"fmt"
	"testing"

	"google.golang.org/protobuf/testing/protocmp"

	"github.com/cosmos/cosmos-sdk/orm/encoding/ormkv"
	"github.com/cosmos/cosmos-sdk/orm/internal/testpb"

	"google.golang.org/protobuf/reflect/protoreflect"
	"gotest.tools/v3/assert"
	"pgregory.net/rapid"

	"github.com/cosmos/cosmos-sdk/orm/internal/testutil"
)

func TestKeyCodec(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		key := testutil.TestKeyCodecGen.Draw(t, "key").(testutil.TestKeyCodec)
		for i := 0; i < 100; i++ {
			keyValues := key.Draw(t, "values")

			bz1 := assertEncDecKey(t, key, keyValues)

			if key.Codec.IsFullyOrdered() {
				// check if ordered keys have ordered encodings
				keyValues2 := key.Draw(t, "values2")
				bz2 := assertEncDecKey(t, key, keyValues2)
				// bytes comparison should equal comparison of values
				assert.Equal(t, key.Codec.CompareValues(keyValues, keyValues2), bytes.Compare(bz1, bz2))
			}
		}
	})
}

func assertEncDecKey(t *rapid.T, key testutil.TestKeyCodec, keyValues []protoreflect.Value) []byte {
	bz, err := key.Codec.Encode(keyValues)
	assert.NilError(t, err)
	keyValues2, err := key.Codec.Decode(bytes.NewReader(bz))
	assert.NilError(t, err)
	key.RequireValuesEqual(t, keyValues, keyValues2)
	return bz
}

func TestPrimaryKeyCodec(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		keyCodec := testutil.TestKeyCodecGen.Draw(t, "keyCodec").(testutil.TestKeyCodec)
		pkCodec := &ormkv.PrimaryKeyCodec{
			KeyCodec: keyCodec.Codec,
			Type:     (&testpb.A{}).ProtoReflect().Type(),
		}
		for i := 0; i < 100; i++ {
			key := keyCodec.Draw(t, fmt.Sprintf("i%d", i))
			pk1 := ormkv.PrimaryKeyEntry{
				Key:   key,
				Value: &testpb.A{},
			}
			k, v, err := pkCodec.EncodeKV(pk1)
			assert.NilError(t, err)
			entry2, err := pkCodec.DecodeKV(k, v)
			assert.NilError(t, err)
			pk2 := entry2.(ormkv.PrimaryKeyEntry)
			assert.Equal(t, 0, pkCodec.CompareValues(pk1.Key, pk2.Key))
			assert.DeepEqual(t, pk1.Value, pk2.Value, protocmp.Transform())
		}
	})
}

func TestUniqueKeyCodec(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		//keyCodec := testutil.TestKeyCodecGen.Draw(t, "keyCodec").(testutil.TestKeyCodec)
		//valueCodec := testutil.TestKeyCodecGen.Draw(t, "valueCodec").(testutil.TestKeyCodec)
		//pkCodec := &ormkv.UniqueKeyCodec{
		//	KeyCodec: keyCodec.Codec,
		//	Type:     (&testpb.A{}).ProtoReflect().Type(),
		//}
		////testKVCodec := testutil.TestPrimaryKeyCodecGen.Draw(t, "primaryKeyCodec").(testutil.TestKVCodec)
		//for i := 0; i < 100; i++ {
		//	key := keyCodec.Draw(t, fmt.Sprintf("i%d", i))
		//	entry := ormkv.PrimaryKeyEntry{
		//		Key:   key,
		//		Value: &testpb.A{},
		//	}
		//	k, v, err := pkCodec.EncodeKV(entry)
		//	assert.NilError(t, err)
		//	entry2, err := pkCodec.DecodeKV(k, v)
		//	assert.NilError(t, err)
		//	assert.DeepEqual(
		//		t, entry, entry2,
		//		cmp.Comparer(func(x, y ormkv.Entry) bool {
		//			pk1 := x.(ormkv.PrimaryKeyEntry)
		//			pk2 := y.(ormkv.PrimaryKeyEntry)
		//			if pkCodec.CompareValues(pk1.Key, pk2.Key) != 0 {
		//				return false
		//			}
		//			return cmp.Equal(pk1.Value, pk2.Value, protocmp.Transform())
		//		}))
		//
		//}
	})
}
