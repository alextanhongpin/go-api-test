package apis

import (
	"github.com/alextanhongpin/go-core-microservice/http/middleware"
	chi "github.com/go-chi/chi/v5"
)

type API struct {
	BearerAuth middleware.Middleware
	*HealthController
}

func (api *API) Register(r chi.Router) {
	r.Get("/health", api.HealthController.Show)
	r.With(api.BearerAuth).Get("/private", api.HealthController.Show)
}
