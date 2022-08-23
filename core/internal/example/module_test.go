package example

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"reflect"
	"testing"
)

func TestModuleContext(t *testing.T) {
	ctxt := sdk.NewContext(nil, tmproto.Header{}, true, nil)
	ctxt = ctxt.WithValue(sdk.SdkContextKey, ctxt)

	modCtxtFactory := sdk.NewModuleContextFactory[ExampleContext]()
	goCtxt := ctxt.Context()
	sdkCtxt := goCtxt.Value(sdk.SdkContextKey)
	require.NotNil(t, sdkCtxt)
	modCtxt := modCtxtFactory.Make(goCtxt)

	modCtxtType := reflect.TypeOf((*ExampleContext)(nil)).Elem()
	require.True(t, reflect.TypeOf(modCtxt).AssignableTo(modCtxtType))
	require.NotPanics(t, func() {
		_ = modCtxt.BlockInfoService()
	})
	require.Panics(t, func() {
		_ = ctxt.GasService()
	})
	// a cast to sdk.Context should fail because capabilities have not been set
	require.Panics(t, func() {
		_ = sdkCtxt.(sdk.Context).EventService()
	})
}
