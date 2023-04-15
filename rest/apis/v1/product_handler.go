package v1

import (
	"context"
	"net/http"

	"github.com/alextanhongpin/go-core-microservice/http/encoding"
	"github.com/alextanhongpin/go-core-microservice/http/types"
	"github.com/go-chi/chi/v5"
)

type ProductUsecase interface {
	Find(ctx context.Context, id string) (*Product, error)
	List(ctx context.Context) ([]Product, error)
}

type ProductHandler struct {
	productUC ProductUsecase
}

func NewProductHandler(productUC ProductUsecase) *ProductHandler {
	return &ProductHandler{
		productUC: productUC,
	}
}

func (h *ProductHandler) Show(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	p, err := h.productUC.Find(r.Context(), id)
	if err != nil {
		encoding.EncodeJSONError(w, err)
		return
	}

	res := types.Result[Product]{
		Data: p,
	}

	encoding.EncodeJSON(w, res, http.StatusOK)
}

func (h *ProductHandler) List(w http.ResponseWriter, r *http.Request) {
	p, err := h.productUC.List(r.Context())
	if err != nil {
		encoding.EncodeJSONError(w, err)
		return
	}

	res := types.Result[[]Product]{
		Data: &p,
	}

	encoding.EncodeJSON(w, res, http.StatusOK)
}
