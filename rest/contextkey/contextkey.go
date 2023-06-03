package contextkey

import (
	"context"
	"errors"
	"fmt"

	"github.com/alextanhongpin/core/http/httputil"
	"github.com/google/uuid"
)

var ErrContextNotFound = errors.New("contextkey not found")

var userIDCtx = httputil.Context[uuid.UUID]("user_id_ctx")

func UserID(ctx context.Context) (uuid.UUID, error) {
	id, ok := userIDCtx.Value(ctx)
	if !ok {
		return uuid.Nil, fmt.Errorf("%w: %s", ErrContextNotFound, userIDCtx)
	}

	return id, nil
}

func SetUserID(ctx context.Context, userID uuid.UUID) context.Context {
	return userIDCtx.WithValue(ctx, userID)
}
