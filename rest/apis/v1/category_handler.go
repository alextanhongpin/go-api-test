package v1

import (
	"errors"
	"net/http"

	"github.com/alextanhongpin/go-core-microservice/http/encoding"
	"github.com/alextanhongpin/go-core-microservice/http/types"
)

type CategoryHandler struct{}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	encoding.EncodeJSONError(w, errors.New("not implemented"))
}

func (h *CategoryHandler) Show(w http.ResponseWriter, r *http.Request) {
	res := types.Result[Category]{
		Data: &Category{
			Name: "Toys",
		},
	}

	encoding.EncodeJSON(w, res, http.StatusOK)
}

func (h *CategoryHandler) List(w http.ResponseWriter, r *http.Request) {
	encoding.EncodeJSONError(w, errors.New("not implemented"))
}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	encoding.EncodeJSONError(w, errors.New("not implemented"))
}

func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	encoding.EncodeJSONError(w, errors.New("not implemented"))
}
