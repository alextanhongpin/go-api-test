package security

import (
	"time"

	"github.com/alextanhongpin/core/http/httputil"
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
	claims := jwt.MapClaims{
		"sub": userID.String(),
	}

	jwtToken, err := httputil.SignJWT(s.secret, claims, duration)

	return jwtToken, err
}
