package v1

import (
	"errors"
	"net/http"

	"github.com/alextanhongpin/go-core-microservice/http/encoding"
	"github.com/alextanhongpin/go-core-microservice/http/types"
)

type CategoryController struct{}

func (h *CategoryController) Create(w http.ResponseWriter, r *http.Request) {
	encoding.EncodeJSONError(w, errors.New("not implemented"))
}

func (h *CategoryController) Show(w http.ResponseWriter, r *http.Request) {
	res := types.Result[Category]{
		Data: &Category{
			Name: "Toys",
		},
	}

	encoding.EncodeJSON(w, res, http.StatusOK)
}

func (h *CategoryController) List(w http.ResponseWriter, r *http.Request) {
	encoding.EncodeJSONError(w, errors.New("not implemented"))
}

func (h *CategoryController) Update(w http.ResponseWriter, r *http.Request) {
	encoding.EncodeJSONError(w, errors.New("not implemented"))
}

func (h *CategoryController) Delete(w http.ResponseWriter, r *http.Request) {
	encoding.EncodeJSONError(w, errors.New("not implemented"))
}
