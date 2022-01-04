package ormlist

import (
	"github.com/cosmos/cosmos-sdk/orm/encoding/encodeutil"
	"github.com/cosmos/cosmos-sdk/orm/internal/listinternal"
)

type Option = listinternal.Option

func Start(values ...interface{}) Option {
	return listinternal.FuncOption(func(options *listinternal.Options) {
		options.Start = encodeutil.ValuesOf(values...)
	})
}

func End(values ...interface{}) Option {
	return listinternal.FuncOption(func(options *listinternal.Options) {
		options.End = encodeutil.ValuesOf(values...)
	})
}

func Prefix(values ...interface{}) Option {
	return listinternal.FuncOption(func(options *listinternal.Options) {
		options.Prefix = encodeutil.ValuesOf(values...)
	})
}

func Reverse() Option {
	return listinternal.FuncOption(func(options *listinternal.Options) {
		options.Reverse = !options.Reverse
	})
}

func Cursor(cursor []byte) Option {
	return listinternal.FuncOption(func(options *listinternal.Options) {
		options.Cursor = cursor
	})
}
