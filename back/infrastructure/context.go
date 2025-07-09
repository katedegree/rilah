package infrastructure

import "context"

type IContext[T any] interface {
	Get(ctx context.Context) (T, bool)
	Set(ctx context.Context, value T) context.Context
}
type ContextKey string
type Context[T any] struct {
	key ContextKey
}

func NewContext[T any](key ContextKey) IContext[T] {
	return &Context[T]{key: key}
}

func (c *Context[T]) Get(ctx context.Context) (T, bool) {
	v := ctx.Value(c.key)
	val, ok := v.(T)
	return val, ok
}

func (c *Context[T]) Set(ctx context.Context, value T) context.Context {
	return context.WithValue(ctx, c.key, value)
}
