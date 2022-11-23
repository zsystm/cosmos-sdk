package valuerenderer

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

type txJsonTest struct {
	Name    string
	Screens []Screen
	Cbor    string
}

func TestLedger(t *testing.T) {
	raw, err := os.ReadFile("../internal/testdata/txs.json")
	require.NoError(t, err)

	var testcases []txJsonTest
	err = json.Unmarshal(raw, &testcases)
	require.NoError(t, err)

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			var buf bytes.Buffer
			err := encode(tc.Screens, &buf)
			require.NoError(t, err)
			fmt.Println(hex.EncodeToString(buf.Bytes()))
			want, err := hex.DecodeString(tc.Cbor)
			require.NoError(t, err)

			require.Equal(t, want, buf.Bytes())
		})
	}
}
