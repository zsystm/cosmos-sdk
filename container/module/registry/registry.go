package registry

import (
	"reflect"

	"google.golang.org/protobuf/proto"

	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/container"
	"github.com/cosmos/cosmos-sdk/container/module/internal"
)

func Resolve(moduleConfig proto.Message) container.Option {
	return ResolveWithModuleName(moduleConfig, string(moduleConfig.ProtoReflect().Descriptor().FullName()))
}

func ResolveWithModuleName(moduleConfig proto.Message, moduleName string) container.Option {
	config, ok := internal.ModuleRegistry[moduleConfig.ProtoReflect().Descriptor().FullName()]
	if !ok {
		return container.Options()
	}

	var opts []container.Option
	for _, provider := range config.Providers {
		opts = append(opts, container.ProvideWithScope(moduleName, provider))
	}

	return container.Options(opts...)
}

func ResolveParam(name protoreflect.FullName, param *internal.ModuleParamType) bool {
	config, ok := internal.ModuleRegistry[name]
	if !ok {
		return false
	}

	*param, ok = config.Params[reflect.TypeOf(*param)]
	if !ok {
		return false
	}

	return true
}
