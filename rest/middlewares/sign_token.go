package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

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
