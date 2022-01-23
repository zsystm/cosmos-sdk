package handler

import (
	"fmt"
	"reflect"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoregistry"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
)

type Resolver interface {
	Resolve(any *anypb.Any, handler interface{}) error
}

type Constructor interface{}

type Config struct {
	HandlerTypes map[string]interface{}
}

type resolver struct {
	typeResolver protoregistry.MessageTypeResolver
	constructors map[reflect.Type]map[protoreflect.FullName]func(interface{}) (reflect.Value, error)
}

func (r resolver) Resolve(any *anypb.Any, handler interface{}) error {
	typ, err := r.typeResolver.FindMessageByURL(any.TypeUrl)
	if err != nil {
		return err
	}

	handlerType := reflect.TypeOf(handler).Elem()
	cmap, ok := r.constructors[handlerType]
	if !ok {
		return fmt.Errorf("unknown handler type %s", handlerType)
	}

	anyName := typ.Descriptor().FullName()
	ctr, ok := cmap[anyName]
	if !ok {
		return fmt.Errorf("no handler of type %s for %s", handlerType, anyName)
	}

	msg := typ.New()
	err = proto.Unmarshal(any.Value, msg.Interface())
	if err != nil {
		return err
	}

	h, err := ctr(msg)
	if err != nil {
		return err
	}

	reflect.ValueOf(handler).Elem().Set(h)

	return nil
}

var _ Resolver = &resolver{}
