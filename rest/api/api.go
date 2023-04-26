package api

import (
	"github.com/alextanhongpin/core/http/middleware"
	chi "github.com/go-chi/chi/v5"
)

type API struct {
	RequireAuth middleware.Middleware
	*AuthController
	*HealthController
}

func (api *API) Register(r chi.Router) {
	// Public routes.
	r.Get("/health", api.HealthController.Show)
	r.Post("/register", api.AuthController.Register)

	// Private routes.
	r.Group(func(r chi.Router) {
		r.Use(api.RequireAuth)
		r.Get("/private", api.HealthController.Show)
	})
}
