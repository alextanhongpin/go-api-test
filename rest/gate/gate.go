package gate

import "context"

func Allow(ctx context.Context, g interface {
	Allow(ctx context.Context) bool
}) bool {
	return g.Allow(ctx)
}
