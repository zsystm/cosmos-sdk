package example

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/capability"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"reflect"
	"testing"
)

func TestModuleContext(t *testing.T) {
	ctxt := sdk.NewContext(nil, tmproto.Header{}, true, nil)
	ctxt = ctxt.WithValue(sdk.SdkContextKey, ctxt)

	modCtxtFactory := capability.NewContextFactory[ModuleContext]()
	sdkCtxt := ctxt.Context().Value(sdk.SdkContextKey)
	require.NotNil(t, sdkCtxt)
	modCtxt := modCtxtFactory.Make(ctxt.Context())

	modCtxtType := reflect.TypeOf((*ModuleContext)(nil)).Elem()
	require.True(t, reflect.TypeOf(modCtxt).AssignableTo(modCtxtType))
	require.NotPanics(t, func() {
		_ = modCtxt.BlockInfoService()
	})
	require.Panics(t, func() {
		_ = ctxt.GasService()
	})
	require.Panics(t, func() {
		_ = ctxt.GasService()
	})
}
