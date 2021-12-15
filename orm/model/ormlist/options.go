package ormlist

import (
	"github.com/cosmos/cosmos-sdk/orm/encoding/encodeutil"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Option interface {
	applyListOption(*opts)
}

type option func(*opts)

func (o option) applyListOption(opts *opts) { o(opts) }

func UseIndex(fieldNames ormtable.FieldNames) Option {
	return option(func(o *opts) {
		o.index = fieldNames
	})
}

func Prefix(values ...interface{}) Option {
	return option(func(o *opts) {
		o.prefix = encodeutil.ValuesOf(values)
	})
}

func Start(values ...interface{}) Option {
	return option(func(o *opts) {
		o.start = encodeutil.ValuesOf(values)
	})
}

func End(values ...interface{}) Option {
	return option(func(o *opts) {
		o.end = encodeutil.ValuesOf(values)
	})
}

func Reverse() Option {
	return option(func(o *opts) {
		o.reverse = true
	})
}

func WithCursor(cursor ormtable.Cursor) Option {
	return option(func(o *opts) {
		o.cursor = cursor
	})
}

//func Filter(filter func(proto.Message) bool) Option {
//	return option(func(o *opts) {
//		o.filter = filter
//	})
//}

type opts struct {
	index   ormtable.FieldNames
	reverse bool
	cursor  ormtable.Cursor
	prefix  []protoreflect.Value
	start   []protoreflect.Value
	end     []protoreflect.Value
	//filter  func(proto.Message) bool
}

func buildOpts(options ...Option) *opts {
	o := &opts{}
	for _, opt := range options {
		opt.applyListOption(o)
	}
	return o
}
