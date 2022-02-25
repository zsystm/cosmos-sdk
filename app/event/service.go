package event

import (
	"context"

	"google.golang.org/protobuf/proto"
)

type Manager interface {
	Emit(proto.Message) error
	EmitSilently(proto.Message)
	EmitLegacy(ctx context.Context, eventType string, attrs ...LegacyEventAttribute) error
}

type LegacyEventAttribute struct {
	Key, Value string
}

func GetManager(ctx context.Context) Manager {
	mgr, ok := ctx.Value(managerContextKey{}).(Manager)
	if !ok {
		return noopEventMgr{}
	}

	return mgr
}

func WithManager(ctx context.Context, manager Manager) context.Context {
	return context.WithValue(ctx, managerContextKey{}, manager)
}

type managerContextKey struct{}

type Hook struct {
	EventType proto.Message
	Func      func(context.Context, proto.Message)
}

type Hooks []Hook

func (Hooks) IsOnePerModuleType() {}

type noopEventMgr struct{}

func (n noopEventMgr) Emit(message proto.Message) error { return nil }

func (n noopEventMgr) EmitSilently(message proto.Message) {}

func (n noopEventMgr) EmitLegacy(ctx context.Context, eventType string, attrs ...LegacyEventAttribute) error {
	return nil
}

var _ Manager = noopEventMgr{}
