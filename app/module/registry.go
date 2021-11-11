package module

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/container"

	"github.com/gogo/protobuf/proto"

	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
)

func Register(initializers ...Initializer) {
	for _, initializer := range initializers {
		typeURL := "/" + proto.MessageName(initializer.Type)
		if _, ok := registry[typeURL]; ok {
			panic(fmt.Errorf("module initializer already defined for %s", typeURL))
		}

		registry[typeURL] = initializer
	}
}

type Initializer struct {
	Type          proto.Message
	RegisterTypes func(codecTypes.InterfaceRegistry)
	Providers     []interface{}
}

var registry = map[string]Initializer{}

func Compose(interfaceRegistry codecTypes.InterfaceRegistry, moduleConfigs map[string]*codecTypes.Any) (container.Option, error) {
	var opts []container.Option
	var initializers []struct {
		name   string
		config *codecTypes.Any
		init   Initializer
	}

	for name, config := range moduleConfigs {
		init, ok := registry[config.TypeUrl]
		if !ok {
			return nil, fmt.Errorf("no module initializer defined for %s", config.TypeUrl)
		}

		interfaceRegistry.RegisterImplementations((*proto.Message)(nil), init.Type)
		init.RegisterTypes(interfaceRegistry)
		initializers = append(initializers, struct {
			name   string
			config *codecTypes.Any
			init   Initializer
		}{name: name, config: config, init: init})
	}

	for _, init := range initializers {
		var config proto.Message
		err := interfaceRegistry.UnpackAny(init.config, &config)
		if err != nil {
			return nil, err
		}

		opts = append(opts, container.Supply(&config))
		opts = append(opts, container.Provide(init.init.Providers...))
	}

	return container.Options(opts...), nil
}
