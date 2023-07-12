package v1

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/alextanhongpin/core/http/response"
	"github.com/alextanhongpin/errcodes/stacktrace"
	"github.com/go-chi/chi/v5"
	"golang.org/x/exp/slog"
)

type ProductUsecase interface {
	Find(ctx context.Context, id string) (*Product, error)
	List(ctx context.Context) ([]Product, error)
}

type ProductController struct {
	productUC ProductUsecase
}

func NewProductController(productUC ProductUsecase) *ProductController {
	return &ProductController{
		productUC: productUC,
	}
}

func (h *ProductController) Show(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	p, err := h.productUC.Find(r.Context(), id)
	if err != nil {
		st := stacktrace.Sprint(err, false)
		fmt.Println(st)
		lines := strings.Split(st, "\n")
		slog.Error("failed to show product", slog.Any("stacktrace", lines))
		response.JSONError(w, err)
		return
	}

	response.JSON(w, response.OK(&p), http.StatusOK)
}

func (h *ProductController) List(w http.ResponseWriter, r *http.Request) {
	p, err := h.productUC.List(r.Context())
	if err != nil {
		response.JSONError(w, err)
		return
	}

	response.JSON(w, response.OK(&p), http.StatusOK)
}
