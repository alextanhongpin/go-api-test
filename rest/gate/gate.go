package gate

import (
	"context"

	"github.com/alextanhongpin/core/http/response"
	"github.com/alextanhongpin/go-api-test/rest/security"
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
	claims, err := security.AuthContext.Value(ctx)
	if err != nil {
		return nil, response.ErrUnauthorized
	}

	sub, err := claims.GetSubject()
	if err != nil {
		return nil, response.ErrUnauthorized
	}

	userID, err := uuid.Parse(sub)
	if err != nil {
		return nil, response.ErrUnauthorized
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
