package example

import (
	"cosmossdk.io/depinject"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"reflect"
	"testing"
)

func TestModuleContext(t *testing.T) {
	ctxt := sdk.NewContext(nil, tmproto.Header{}, true, nil)
	modCtxtFactory := sdk.NewModuleContextFactory[ModuleContext](depinject.ModuleKey{})
	modCtxt := modCtxtFactory.Make(ctxt)

	modCtxtType := reflect.TypeOf((*ModuleContext)(nil)).Elem()
	require.True(t, reflect.TypeOf(modCtxt).AssignableTo(modCtxtType))
}
