package event

import (
	"context"

	"google.golang.org/protobuf/proto"
)

type Service interface {
}

func Emit(context.Context, proto.Message) error {
	return nil
}

func EmitSilently(context.Context, proto.Message) error {
	return nil
}

func EmitLegacy(ctx context.Context, eventType string, attrs ...LegacyEventAttribute) error {
	return nil
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
