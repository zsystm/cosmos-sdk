package valuerenderer_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"cosmossdk.io/tx/textual/valuerenderer"
	"github.com/stretchr/testify/require"

	bankv1beta1 "cosmossdk.io/api/cosmos/bank/v1beta1"
	"cosmossdk.io/tx/textual/internal/testpb"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type repeatedJsonTest struct {
	Proto    *testpb.Baz
	Metadata map[string]*bankv1beta1.Metadata
	Screens  []valuerenderer.Screen
}

func TestRepeatedJsonTestcases(t *testing.T) {
	raw, err := os.ReadFile("../internal/testdata/repeated.json")
	require.NoError(t, err)

	var testcases []repeatedJsonTest
	err = json.Unmarshal(raw, &testcases)
	require.NoError(t, err)

	tr := valuerenderer.NewTextual(mockCoinMetadataQuerier)
	for i, tc := range testcases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			// Create a context.Context containing all coins metadata, to simulate
			// that they are in state.
			ctx := context.Background()
			for _, coin := range tc.Proto.Price {
				ctx = context.WithValue(ctx, mockCoinMetadataKey(coin.Denom), tc.Metadata[coin.Denom])
			}

			rend := valuerenderer.NewMessageValueRenderer(&tr, (&testpb.Baz{}).ProtoReflect().Descriptor())
			require.NoError(t, err)

			screens, err := rend.Format(ctx, protoreflect.ValueOf(tc.Proto.ProtoReflect()))
			require.NoError(t, err)
			require.Equal(t, tc.Screens, screens)

			//val, err := rend.Parse(context.Background(), screens)
			//require.NoError(t, err)
			//msg := val.Message().Interface()
			//require.IsType(t, &testpb.Baz{}, msg)
			//foo := msg.(*testpb.Baz)
			//require.True(t, proto.Equal(foo, tc.Proto))
		})
	}
}
