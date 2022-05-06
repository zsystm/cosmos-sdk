package module

import (
	"reflect"

	"google.golang.org/protobuf/proto"

	"github.com/cosmos/cosmos-sdk/core/internal"
)

func Register(msg proto.Message, options ...Option) {
	init := &internal.ModuleInitializer{
		ConfigProtoType: msg.ProtoReflect().Type(),
		ConfigGoType:    reflect.TypeOf(msg),
	}

	internal.ModuleRegistry[msg.ProtoReflect().Descriptor().FullName()] = init

	for _, option := range options {
		init.Error = option.apply(init)
		if init.Error != nil {
			return
		}
	}
}
