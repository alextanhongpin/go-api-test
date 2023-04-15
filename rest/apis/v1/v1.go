package v1

import "github.com/go-chi/chi/v5"

type API struct {
	*CategoryHandler
	*ProductHandler
}

func (api *API) Register(r chi.Router) {
	r.Route("/v1", func(r chi.Router) {
		r.Route("/categories", func(r chi.Router) {
			h := api.CategoryHandler

			r.Get("/", h.Show)
		})

		r.Route("/products", func(r chi.Router) {
			h := api.ProductHandler

			r.Get("/{id}", h.Show)
			r.Get("/", h.List)
		})
	})

}
