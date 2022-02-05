package event

import (
	"context"

	"google.golang.org/protobuf/proto"
)

type Service interface {
	Emit(context.Context, proto.Message) error
	EmitSilently(context.Context, proto.Message) error
	EmitLegacy(ctx context.Context, eventType string, attrs ...LegacyEventAttribute) error
}

type LegacyEventAttribute struct {
	Key, Value string
}

type Hook struct {
	EventType proto.Message
	Func      func(context.Context, proto.Message)
}

type Hooks []Hook

func (Hooks) IsOnePerModuleType() {}
