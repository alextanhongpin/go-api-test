package v1

import "github.com/go-chi/chi/v5"

type API struct {
	*CategoryHandler
	*ProductHandler
}

func (api *API) Register(r chi.Router) {
	r.Route("/v1", func(r chi.Router) {
		r.Route("/categories", func(r chi.Router) {
			r.Get("/{id}", api.CategoryHandler.Show)
			r.Get("/", api.CategoryHandler.List)
			r.Post("/", api.CategoryHandler.Create)
			r.Patch("/", api.CategoryHandler.Update)
			r.Delete("/", api.CategoryHandler.Delete)
		})

		r.Route("/products", func(r chi.Router) {
			r.Get("/{id}", api.ProductHandler.Show)
			r.Get("/", api.ProductHandler.List)
		})
	})
}
