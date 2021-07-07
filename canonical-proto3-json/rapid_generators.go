package canonical_proto3_json

import (
	"time"

	cosmostypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/gogo/protobuf/types"
	"pgregory.net/rapid"
)

var (
	genCompleteTest = rapid.Custom(func(t *rapid.T) *CompleteTest {
		return &CompleteTest{
			NumericalTest:       &CompleteTest_NumericalTest{},
			TestEnum:            0,
			TestMap:             map[int32]int64{},
			TestTimeMap:         map[int32]*types.Duration{},
			TestRepeatedFixed32: []uint32{},
			TestBool:            rapid.Bool().Draw(t, "TestBool").(bool),
			TestString:          "",
			TestBytes:           []byte{},
			TestAny:             &cosmostypes.Any{},
			TestTimestamp:       &types.Timestamp{},
			TestStdTimestamp:    &time.Time{},
			TestDuration:        &types.Duration{},
			TestStruct:          genStruct.Draw(t, "TestStruct").(*types.Struct),
			TestBoolValue:       &types.BoolValue{},
			TestBytesValue:      &types.BytesValue{},
			TestDoubleValue:     &types.DoubleValue{},
			TestFloatValue:      &types.FloatValue{},
			TestInt32Value:      &types.Int32Value{},
			TestInt64Value:      &types.Int64Value{},
			TestStringValue:     &types.StringValue{},
			TestUint32Value:     &types.UInt32Value{},
			TestUint64Value:     &types.UInt64Value{},
			TestFieldMask:       &types.FieldMask{},
			TestEmpty:           &types.Empty{},
		}
	})

	genStruct = rapid.Custom(func(t *rapid.T) *types.Struct {
		fields := rapid.MapOfN(rapid.String(), genValue, 0, 10).Draw(t, "fields").(map[string]*types.Value)

		return &types.Struct{
			Fields: fields,
		}
	})

	genValue = rapid.Custom(func(t *rapid.T) *types.Value {
		return rapid.Just(&types.Value{
			Kind: &types.Value_BoolValue{
				BoolValue: true,
			},
		}).Draw(t, "value").(*types.Value)
	})
)
