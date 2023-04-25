package apis

import (
	"github.com/alextanhongpin/core/http/middleware"
	chi "github.com/go-chi/chi/v5"
)

type API struct {
	BearerAuth middleware.Middleware
	*AuthController
	*HealthController
}

func (api *API) Register(r chi.Router) {
	r.Get("/health", api.HealthController.Show)
	r.With(api.BearerAuth).Get("/private", api.HealthController.Show)

	r.Post("/register", api.AuthController.Register)
}
