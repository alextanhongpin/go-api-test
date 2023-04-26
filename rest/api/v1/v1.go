package v1

import (
	"github.com/alextanhongpin/core/http/middleware"
	"github.com/go-chi/chi/v5"
)

type API struct {
	RequireAuth middleware.Middleware
	*CategoryController
	*ProductController
}

func (api *API) Register(r chi.Router) {
	r.Route("/v1", func(r chi.Router) {
		r.Route("/categories", func(r chi.Router) {
			r.Get("/{id}", api.CategoryController.Show)
			r.Get("/", api.CategoryController.List)
			r.With(api.RequireAuth).Post("/", api.CategoryController.Create)
			r.Patch("/", api.CategoryController.Update)
			r.Delete("/", api.CategoryController.Delete)
		})

		r.Route("/products", func(r chi.Router) {
			r.Get("/{id}", api.ProductController.Show)
			r.Get("/", api.ProductController.List)
		})

		r.Group(func(r chi.Router) {
			r.Use(api.RequireAuth)
		})
	})
}
