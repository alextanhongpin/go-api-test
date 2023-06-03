package gate

import (
	"context"
	"fmt"

	"github.com/alextanhongpin/core/http/response"
	"github.com/alextanhongpin/go-api-test/rest/contextkey"
	"github.com/google/uuid"
)

type GateKeeper interface {
	Allow(user User) bool
}

type User interface {
	ID() uuid.UUID
	//Role() string
	//Scopes() []string
}

type user struct {
	id uuid.UUID
}

func (u *user) ID() uuid.UUID {
	return u.id
}

type Gate struct {
	user User
}

func New(ctx context.Context) (*Gate, error) {
	userID, err := contextkey.UserID(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", response.ErrUnauthorized, err)
	}

	return &Gate{
		user: &user{id: userID},
	}, nil
}

func (g *Gate) User() User {
	return g.user
}

func (g *Gate) Allow(gatekeepers ...GateKeeper) bool {
	for _, gks := range gatekeepers {
		if !gks.Allow(g.user) {
			return false
		}
	}

	return g.user.ID() != uuid.Nil
}
