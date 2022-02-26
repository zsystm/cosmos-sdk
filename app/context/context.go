package context

import (
	"context"

	internalcontext "github.com/cosmos/cosmos-sdk/app/internal/context"
)

func NewContext(baseCtx context.Context) context.Context {
	if baseCtx == nil {
		baseCtx = context.Background()
	}

	return internalcontext.Context{BaseCtx: baseCtx}
}

func WithValue(ctx context.Context, key, value interface{}) context.Context {
	if c, ok := ctx.(interface {
		WithValue(key, value interface{}) context.Context
	}); ok {
		return c.WithValue(key, value)
	}

	return context.WithValue(ctx, key, value)
}
