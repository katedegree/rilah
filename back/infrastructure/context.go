package infrastructure

import "context"

type IContext[T any] interface {
	Get(ctx context.Context) T
	Set(ctx context.Context, value T) context.Context
}
type ContextKey string
type Context[T any] struct {
	key ContextKey
}

func NewContext[T any](key ContextKey) IContext[T] {
	return &Context[T]{key: key}
}

func (c *Context[T]) Get(ctx context.Context) T {
	return ctx.Value(c.key).(T)
}

func (c *Context[T]) Set(ctx context.Context, value T) context.Context {
	return context.WithValue(ctx, c.key, value)
}
