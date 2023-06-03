package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/alextanhongpin/core/http/httputil"
	"github.com/alextanhongpin/core/http/response"
	"github.com/alextanhongpin/go-api-test/rest/contextkey"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func RequireAuth(secret []byte) Middleware {
	fn := func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {

			token, ok := httputil.BearerAuth(r)
			if ok {
				claims, err := httputil.VerifyJWT(secret, token)
				if err != nil {
					if errors.Is(err, jwt.ErrTokenExpired) {
						err = fmt.Errorf("%w: %w", response.ErrUnauthorized, err)
					}
					response.JSONError(w, err)
					return
				}

				ctx := r.Context()
				sub, err := claims.GetSubject()
				if err != nil {
					response.JSONError(w, err)
					return
				}
				userID, err := uuid.Parse(sub)
				if err != nil {
					response.JSONError(w, err)
					return
				}

				ctx = contextkey.SetUserID(ctx, userID)
				r = r.WithContext(ctx)
			}

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}

	return fn
}
