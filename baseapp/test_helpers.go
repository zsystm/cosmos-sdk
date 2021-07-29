package baseapp

import (
	"context"
	"fmt"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"google.golang.org/grpc"

	"github.com/cosmos/cosmos-sdk/container"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (app *BaseApp) Check(txEncoder sdk.TxEncoder, tx sdk.Tx) (sdk.GasInfo, *sdk.Result, error) {
	// runTx expects tx bytes as argument, so we encode the tx argument into
	// bytes. Note that runTx will actually decode those bytes again. But since
	// this helper is only used in tests/simulation, it's fine.
	bz, err := txEncoder(tx)
	if err != nil {
		return sdk.GasInfo{}, nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "%s", err)
	}
	return app.runTx(runTxModeCheck, bz)
}

func (app *BaseApp) Simulate(txBytes []byte) (sdk.GasInfo, *sdk.Result, error) {
	return app.runTx(runTxModeSimulate, txBytes)
}

func (app *BaseApp) Deliver(txEncoder sdk.TxEncoder, tx sdk.Tx) (sdk.GasInfo, *sdk.Result, error) {
	// See comment for Check().
	bz, err := txEncoder(tx)
	if err != nil {
		return sdk.GasInfo{}, nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "%s", err)
	}
	return app.runTx(runTxModeDeliver, bz)
}

// Context with current {check, deliver}State of the app used by tests.
func (app *BaseApp) NewContext(isCheckTx bool, header tmproto.Header) sdk.Context {
	if isCheckTx {
		return sdk.NewContext(app.checkState.ms, header, true, app.logger).
			WithMinGasPrices(app.minGasPrices)
	}

	return sdk.NewContext(app.deliverState.ms, header, false, app.logger)
}

func (app *BaseApp) NewUncachedContext(isCheckTx bool, header tmproto.Header) sdk.Context {
	return sdk.NewContext(app.cms, header, isCheckTx, app.logger)
}

var UnitTestFixture = container.Provide(
	func(baseApp *BaseApp) (
		func() sdk.Context,
		func() context.Context) {
		getSDKContext := func() sdk.Context { return baseApp.deliverState.ctx }
		getContext := func() context.Context {
			ctx := getSDKContext()
			return sdk.WrapSDKContext(ctx)
		}
		return getSDKContext, getContext
	})

type unitTestClient struct {
	baseApp *BaseApp
}

func (u unitTestClient) Invoke(ctx context.Context, _ string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
	msg, ok := args.(sdk.Msg)
	if !ok {
		return fmt.Errorf("expected instance of %T, got %T", msg, args)
	}

	err := msg.ValidateBasic()
	if err != nil {
		return err
	}

	if handler := u.baseApp.msgServiceRouter.Handler(msg); handler != nil {
		res, err := handler(sdk.UnwrapSDKContext(ctx), msg)
		if err != nil {
			return err
		}
		reply = res
	}
	panic("TODO")
}

func (u unitTestClient) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	return nil, fmt.Errorf("not supported")
}

var _ grpc.ClientConnInterface = &unitTestClient{}
