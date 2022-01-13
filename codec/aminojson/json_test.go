package aminojson

import (
	"testing"

	"github.com/stretchr/testify/require"

	bankv1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/bank/v1beta1"
	basev1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
)

func Test1(t *testing.T) {
	bz, err := Marshal(&bankv1beta1.MsgSend{
		FromAddress: "foo213325",
		ToAddress:   "foo32t5sdfh",
		Amount: []*basev1beta1.Coin{
			{
				Denom:  "bar",
				Amount: "1234",
			},
		},
	})
	require.NoError(t, err)
	require.Equal(t, "", string(bz))
}
