package v1_test

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	tu "github.com/alextanhongpin/go-api-test/internal/testutils"
	"github.com/alextanhongpin/go-api-test/mocks"
	v1 "github.com/alextanhongpin/go-api-test/rest/apis/v1"
	chi "github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestProductHandlerShow(t *testing.T) {
	tests := []struct {
		name    string
		find    *v1.Product
		findErr error
		want    []byte
		status  int
	}{
		{
			name:   "success",
			find:   &v1.Product{Name: "Colorful Socks"},
			status: http.StatusOK,
			want:   []byte(`{"data": {"name": "Colorful Socks"}}`),
		},
		{
			name:    "failed",
			findErr: errors.New("db error"),
			status:  http.StatusInternalServerError,
			want:    tu.InternalServerErrorBytes,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert := assert.New(t)
			uc := new(mocks.ProductUsecase)
			uc.On("Find", tu.ContextType, tu.StringType).Return(tc.find, tc.findErr).Once()

			handler := v1.NewProductHandler(uc)

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/v1/products/colorful-socks", nil)

			// We need to inject the URL params manually when using chi router.
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", "1")
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

			handler.Show(w, r)
			res := w.Result()
			defer res.Body.Close()

			assert.Equal(tc.status, res.StatusCode, "status code does not match")
			got, err := ioutil.ReadAll(res.Body)
			assert.Nil(err)
			tu.CmpJSON(t, tc.want, got)
		})
	}
}

func TestProductHandlerList(t *testing.T) {
	tests := []struct {
		name    string
		list    []v1.Product
		listErr error
		want    []byte
		status  int
	}{
		{
			name: "success",
			list: []v1.Product{
				{Name: "red socks"},
				{Name: "green socks"},
				{Name: "blue socks"},
			},
			status: http.StatusOK,
			want: []byte(`{
	"data": [
		{"name": "red socks"},
		{"name": "green socks"},
		{"name": "blue socks"}
	]
}`),
		},
		{
			name:    "failed",
			listErr: errors.New("db error"),
			status:  http.StatusInternalServerError,
			want:    tu.InternalServerErrorBytes,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert := assert.New(t)

			uc := new(mocks.ProductUsecase)
			uc.On("List", tu.ContextType).Return(tc.list, tc.listErr).Once()

			handler := v1.NewProductHandler(uc)
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/v1/products", nil)

			handler.List(w, r)
			res := w.Result()
			defer res.Body.Close()

			assert.Equal(tc.status, res.StatusCode, "status code does not match")
			got, err := ioutil.ReadAll(res.Body)
			assert.Nil(err)
			tu.CmpJSON(t, tc.want, got)
		})
	}
}
