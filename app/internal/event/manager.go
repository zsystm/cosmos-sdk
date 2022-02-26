package event

import (
	"google.golang.org/protobuf/proto"
)

type Manager interface {
	Emit(proto.Message) error
	EmitSilently(proto.Message)
	EmitLegacy(eventType string, attrs ...LegacyEventAttribute) error
}

type LegacyEventAttribute struct {
	Key, Value string
}
