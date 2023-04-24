package apis

import (
	"net/http"
	"time"

	"github.com/alextanhongpin/go-api-test/rest/middlewares"
	"github.com/alextanhongpin/go-core-microservice/http/encoding"
	"github.com/alextanhongpin/go-core-microservice/http/types"
	"github.com/google/uuid"
)

type AuthController struct {
	TokenSigner *middlewares.TokenSigner
}

type RegisterResponse struct {
	AccessToken string  `json:"accessToken"`
	ExpiresIn   float64 `json:"expiresIn"`
}

func (h *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	fakeUserID := uuid.New()
	duration := 1 * time.Hour
	token, err := h.TokenSigner.SignUserID(fakeUserID, duration)
	if err != nil {
		encoding.EncodeJSONError(w, types.ErrInternal)
		return
	}

	res := types.Result[RegisterResponse]{
		Data: &RegisterResponse{
			AccessToken: token,
			ExpiresIn:   duration.Seconds(),
		},
	}

	encoding.EncodeJSON(w, res, http.StatusOK)
}
