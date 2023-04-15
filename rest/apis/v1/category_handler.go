package v1

import (
	"net/http"

	"github.com/alextanhongpin/go-core-microservice/http/encoding"
	"github.com/alextanhongpin/go-core-microservice/http/types"
)

type CategoryHandler struct{}

func (h *CategoryHandler) Show(w http.ResponseWriter, r *http.Request) {
	res := types.Result[Category]{
		Data: &Category{
			Name: "Toys",
		},
	}

	encoding.EncodeJSON(w, res, http.StatusOK)
}
