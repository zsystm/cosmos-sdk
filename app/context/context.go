package context

import (
	"context"
	"time"
)

type Context struct {
	baseCtx context.Context
	values  map[interface{}]interface{}
}

func NewContext(baseCtx context.Context) *Context {
	return &Context{baseCtx: baseCtx, values: map[interface{}]interface{}{}}
}

func (c *Context) Deadline() (deadline time.Time, ok bool) {
	if c.baseCtx == nil {
		return time.Time{}, false
	}

	return c.baseCtx.Deadline()
}

func (c *Context) Done() <-chan struct{} {
	if c.baseCtx == nil {
		return nil
	}

	return c.baseCtx.Done()
}

func (c *Context) Err() error {
	if c.baseCtx == nil {
		return nil
	}

	return c.baseCtx.Err()
}

func (c *Context) Value(key interface{}) interface{} {
	v, ok := c.values[key]
	if ok {
		return v
	}

	if c.baseCtx != nil {
		return c.baseCtx.Value(key)
	}

	return nil
}

// NOTE: this mutates the underlying context and even though it's faster is maybe not what we want
// because we want to be able to have contexts inherit
func (c *Context) WithValue(key, value interface{}) context.Context {
	if c.values == nil {
		c.values = map[interface{}]interface{}{}
	}

	c.values[key] = value
	return c
}

func WithValue(ctx context.Context, key, value interface{}) context.Context {
	if c, ok := ctx.(*Context); ok {
		return c.WithValue(key, value)
	}

	return context.WithValue(ctx, key, value)
}

var _ context.Context = &Context{}
