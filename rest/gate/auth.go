package gate

import (
	"context"

	"github.com/alextanhongpin/go-core-microservice/http/middleware"
	"github.com/google/uuid"
)

type User struct {
	ID  uuid.UUID
	Err error
}

func (u *User) Allow(ctx context.Context) bool {
	u.Err = u.checkIsLoggedIn(ctx)
	return u.Err == nil
}

func (u *User) checkIsLoggedIn(ctx context.Context) error {
	claims, err := middleware.AuthContext.Value(ctx)
	if err != nil {
		return err
	}

	userIDStr, err := claims.GetSubject()
	if err != nil {
		return err
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return err
	}

	u.ID = userID

	return nil
}
