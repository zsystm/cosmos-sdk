package valuerenderer_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"cosmossdk.io/tx/textual/valuerenderer"
	"github.com/stretchr/testify/require"

	"cosmossdk.io/tx/textual/internal/testpb"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type repeatedJsonTest struct {
	Proto   *testpb.Baz
	Screens []valuerenderer.Screen
}

func TestRepeatedJsonTestcases(t *testing.T) {
	raw, err := os.ReadFile("../internal/testdata/repeated.json")
	require.NoError(t, err)

	var testcases []repeatedJsonTest
	err = json.Unmarshal(raw, &testcases)
	require.NoError(t, err)

	tr := valuerenderer.NewTextual(EmptyCoinMetadataQuerier)
	for i, tc := range testcases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			rend := valuerenderer.NewRepeatedValueRenderer(&tr, (&testpb.Baz{}).ProtoReflect().Descriptor(), valuerenderer.NewMessageValueRenderer(&tr, (&testpb.Foo{}).ProtoReflect().Descriptor()))

			screens, err := rend.Format(context.Background(), protoreflect.ValueOfList(NewGenericList(tc.Proto.Messages)))
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
