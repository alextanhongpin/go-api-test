package middlewares

import (
	"net/http"

	"github.com/alextanhongpin/go-api-test/rest/contextkey"
	"github.com/alextanhongpin/go-core-microservice/http/encoding"
	"github.com/alextanhongpin/go-core-microservice/http/middleware"
	"github.com/alextanhongpin/go-core-microservice/http/types"
	"github.com/google/uuid"
)

func MustUserID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		claims, err := middleware.AuthContext.Value(ctx)
		if err != nil {
			encoding.EncodeJSONError(w, types.ErrUnauthorized)
			return
		}

		userIDStr, err := claims.GetSubject()
		if err != nil {
			encoding.EncodeJSONError(w, types.ErrUnauthorized)
			return
		}

		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			encoding.EncodeJSONError(w, types.ErrInternal)
			return
		}

		ctx = contextkey.UserID.WithValue(ctx, userID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
