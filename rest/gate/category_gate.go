package gate

import (
	"context"

	"github.com/google/uuid"
)

type CategoryCreator struct {
	UserID uuid.UUID
}

func (c *CategoryCreator) Allow(ctx context.Context) bool {
	return false
}
