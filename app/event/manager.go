package event

import (
	"context"

	appcontext "github.com/cosmos/cosmos-sdk/app/context"
	internalcontext "github.com/cosmos/cosmos-sdk/app/internal/context"
	"github.com/cosmos/cosmos-sdk/app/internal/event"

	"google.golang.org/protobuf/proto"
)

type (
	Manager              = event.Manager
	LegacyEventAttribute = event.LegacyEventAttribute
)

func GetManager(ctx context.Context) Manager {
	sdkCtx, ok := ctx.(internalcontext.Context)
	if ok {
		return sdkCtx.EventManager
	}

	if mgr, ok := ctx.Value(managerContextKey{}).(Manager); ok {
		return mgr
	}

	return noopEventMgrInstance
}

func WithManager(ctx context.Context, manager Manager) context.Context {
	sdkCtx, ok := ctx.(internalcontext.Context)
	if ok {
		sdkCtx.EventManager = manager
		return sdkCtx
	}

	return appcontext.WithValue(ctx, managerContextKey{}, manager)
}

type managerContextKey struct{}

var noopEventMgrInstance = noopEventMgr{}

type noopEventMgr struct{}

func (n noopEventMgr) Emit(message proto.Message) error { return nil }

func (n noopEventMgr) EmitSilently(message proto.Message) {}

func (n noopEventMgr) EmitLegacy(eventType string, attrs ...LegacyEventAttribute) error {
	return nil
}

var _ Manager = noopEventMgr{}
