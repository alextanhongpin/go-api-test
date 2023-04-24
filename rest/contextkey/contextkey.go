package contextkey

import (
	"context"
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("contextkey: key not found")

type ContextKey[T any] string

func (key ContextKey[T]) Value(ctx context.Context) (T, bool) {
	v, ok := ctx.Value(key).(T)
	return v, ok
}

func (key ContextKey[T]) MustValue(ctx context.Context) T {
	v, ok := key.Value(ctx)
	if !ok {
		panic(fmt.Errorf("%w: %v", ErrNotFound, key))
	}

	return v
}

func (key ContextKey[T]) WithValue(ctx context.Context, val T) context.Context {
	ctx = context.WithValue(ctx, key, val)
	return ctx
}
