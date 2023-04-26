package security

import (
	"context"
	"time"

	"github.com/alextanhongpin/core/http/security"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func UserIDFromContext(ctx context.Context) (uuid.UUID, error) {
	claims, err := security.AuthContext.Value(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	sub, err := claims.GetSubject()
	if err != nil {
		return uuid.Nil, err
	}

	userID, err := uuid.Parse(sub)
	if err != nil {
		return uuid.Nil, err
	}

	return userID, nil
}

type TokenSigner struct {
	secret []byte
}

func NewTokenSigner(secret []byte) *TokenSigner {
	return &TokenSigner{
		secret: secret,
	}
}

func (s *TokenSigner) SignUserID(userID uuid.UUID, duration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID.String(),
		"exp": time.Now().Add(duration).Unix(),
	})

	jwtToken, err := token.SignedString(s.secret)
	return jwtToken, err
}
