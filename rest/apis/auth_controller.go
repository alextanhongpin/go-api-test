package apis

import (
	"net/http"
	"time"

	"github.com/alextanhongpin/core/http/response"
	"github.com/alextanhongpin/go-api-test/rest/security"
	"github.com/google/uuid"
)

type AuthController struct {
	TokenSigner *security.TokenSigner
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
		response.JSONError(w, response.ErrInternal)
		return
	}

	response.JSON(w, response.OK(&RegisterResponse{
		AccessToken: token,
		ExpiresIn:   duration.Seconds(),
	}), http.StatusOK)
}
