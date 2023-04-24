package v1

import (
	"github.com/alextanhongpin/go-api-test/rest/middlewares"
	"github.com/alextanhongpin/go-core-microservice/http/middleware"
	"github.com/go-chi/chi/v5"
)

type API struct {
	BearerAuth middleware.Middleware
	*CategoryController
	*ProductController
}

func (api *API) Register(r chi.Router) {
	r.Route("/v1", func(r chi.Router) {
		r.Route("/categories", func(r chi.Router) {
			r.Get("/{id}", api.CategoryController.Show)
			r.Get("/", api.CategoryController.List)
			r.With(api.BearerAuth, middlewares.MustUserID).Post("/", api.CategoryController.Create)
			r.Patch("/", api.CategoryController.Update)
			r.Delete("/", api.CategoryController.Delete)
		})

		r.Route("/products", func(r chi.Router) {
			r.Get("/{id}", api.ProductController.Show)
			r.Get("/", api.ProductController.List)
		})
	})
}
