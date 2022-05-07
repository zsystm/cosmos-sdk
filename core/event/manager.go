package event

import "google.golang.org/protobuf/proto"

// Manager represents an event manager.
type Manager interface {
	Emit(proto.Message) error
	EmitLegacy(eventType string, attrs ...LegacyEventAttribute) error
}

// LegacyEventAttribute is a legacy (untyped) event attribute.
type LegacyEventAttribute struct {
	Key, Value string
}
