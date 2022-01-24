package module

import (
	"google.golang.org/protobuf/proto"

	"github.com/cosmos/cosmos-sdk/app/internal"
)

type Option = internal.Option

func RegisterModule(message proto.Message, options ...Option) {
	init := &internal.Initializer{
		Type:      message.ProtoReflect().Type(),
		Providers: nil,
	}

	// TODO
	//for _, option := range options {
	//	option.apply(init)
	//}

	internal.Registry[message.ProtoReflect().Descriptor().FullName()] = init
}

func Declare(interface{}) Option {
	panic("TODO")
}

func Provide(constructors ...interface{}) Option {
	panic("TODO")
}

//func Compose(moduleConfigs map[string]*anypb.Any, resolver protoregistry.MessageTypeResolver) (container.Option, error) {
//	var opts []container.Option
//	var initializers []struct {
//		name   string
//		config *anypb.Any
//		init   *initializer
//	}
//
//	for name, config := range moduleConfigs {
//		typ :=
//		init, ok := registry[config.TypeUrl]
//		if !ok {
//			return nil, fmt.Errorf("no module initializer defined for %s", config.TypeUrl)
//		}
//
//		interfaceRegistry.RegisterImplementations((*proto.Message)(nil), init.Type)
//		init.RegisterTypes(interfaceRegistry)
//		initializers = append(initializers, struct {
//			name   string
//			config *codecTypes.Any
//			init   Initializer
//		}{name: name, config: config, init: init})
//	}
//
//	for _, init := range initializers {
//		var config proto.Message
//		err := interfaceRegistry.UnpackAny(init.config, &config)
//		if err != nil {
//			return nil, err
//		}
//
//		opts = append(opts, container.Supply(&config))
//		opts = append(opts, container.Provide(init.init.Providers...))
//	}
//
//	return container.Options(opts...), nil
//}
