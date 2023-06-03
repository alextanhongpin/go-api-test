package v1_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alextanhongpin/core/test/testutil"
	"github.com/alextanhongpin/go-api-test/mocks"
	v1 "github.com/alextanhongpin/go-api-test/rest/api/v1"
	chi "github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
)

func TestProductControllerShow(t *testing.T) {
	tests := []struct {
		name       string
		find       *v1.Product
		findErr    error
		statusCode int
	}{
		{
			name:       "success",
			find:       &v1.Product{Name: "Colorful Socks"},
			statusCode: http.StatusOK,
		},
		{
			name:       "failed",
			findErr:    errors.New("db error"),
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			uc := new(mocks.ProductUsecase)
			uc.On("Find", mock.Anything, mock.Anything).Return(tc.find, tc.findErr).Once()

			handler := v1.NewProductController(uc).Show

			// We need to inject the URL params manually when using chi router.
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", "1")

			r := httptest.NewRequest("GET", "/v1/products/colorful-socks", nil)
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
			testutil.DumpHTTP(t, r, handler)
		})
	}
}

func TestProductControllerList(t *testing.T) {
	tests := []struct {
		name       string
		list       []v1.Product
		listErr    error
		statusCode int
	}{
		{
			name: "success",
			list: []v1.Product{
				{Name: "red socks"},
				{Name: "green socks"},
				{Name: "blue socks"},
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "failed",
			listErr:    errors.New("db error"),
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			uc := new(mocks.ProductUsecase)
			uc.On("List", mock.Anything).Return(tc.list, tc.listErr).Once()

			handler := v1.NewProductController(uc).List
			r := httptest.NewRequest("GET", "/v1/products", nil)
			testutil.DumpHTTP(t, r, handler)
		})
	}
}
