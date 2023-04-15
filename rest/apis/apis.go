package apis

import (
	"github.com/alextanhongpin/go-core-microservice/http/middleware"
	chi "github.com/go-chi/chi/v5"
)

type API struct {
	*HealthHandler
	BearerMW middleware.Middleware
}

func (api *API) Register(r chi.Router) {
	r.Get("/health", api.HealthHandler.Show)
	r.With(api.BearerMW).Get("/private", api.HealthHandler.Show)
}
