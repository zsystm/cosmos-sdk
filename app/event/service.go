package event

import (
	"context"

	"google.golang.org/protobuf/proto"
)

type Service struct {
}

func (s *Service) Emit(context.Context, proto.Message) error {
	panic("TODO")
}

func (s *Service) EmitSilently(context.Context, proto.Message) error {
	panic("TODO")
}

func (s *Service) EmitLegacy(ctx context.Context, eventType string, attrs ...LegacyEventAttribute) error {
	panic("TODO")
}

type LegacyEventAttribute struct {
	Key, Value string
}
