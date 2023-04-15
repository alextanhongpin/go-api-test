package v1_test

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alextanhongpin/go-api-test/internal/testutils"
	"github.com/alextanhongpin/go-api-test/mocks"
	v1 "github.com/alextanhongpin/go-api-test/rest/apis/v1"
	chi "github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProductHandlerShow(t *testing.T) {
	uc := new(mocks.ProductUsecase)

	tests := []struct {
		name       string
		setup      func()
		statusCode int
		res        []byte
	}{
		{
			name: "success",
			setup: func() {
				pdt := &v1.Product{Name: "Colorful Socks"}
				uc.On("Find", testutils.OfTypeContext, mock.AnythingOfType("string")).Return(pdt, nil).Once()
			},
			statusCode: http.StatusOK,
			res: []byte(`{
	"data": {
		"name": "Colorful Socks"
	}
}`),
		},
		{
			name: "failed",
			setup: func() {
				uc.On("Find", testutils.OfTypeContext, mock.AnythingOfType("string")).Return(nil, errors.New("db error")).Once()
			},
			statusCode: http.StatusInternalServerError,
			res:        testutils.InternalServerErrorBytes,
		},
	}

	for _, tc := range tests {
		tc := tc
		tc.setup()

		t.Run(tc.name, func(t *testing.T) {

			assert := assert.New(t)

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

			assert.Equal(tc.statusCode, res.StatusCode, "status code does not match")
			got, err := ioutil.ReadAll(res.Body)
			assert.Nil(err)
			testutils.CmpJSON(t, tc.res, got)
		})
	}
}

func TestProductHandlerList(t *testing.T) {
	uc := new(mocks.ProductUsecase)

	tests := []struct {
		name       string
		setup      func()
		res        []byte
		statusCode int
	}{
		{
			name: "success",
			setup: func() {
				ret := []v1.Product{
					{Name: "red socks"},
					{Name: "green socks"},
					{Name: "blue socks"},
				}
				uc.On("List", testutils.OfTypeContext).Return(ret, nil).Once()
			},
			statusCode: http.StatusOK,
			res: []byte(`{
	"data": [
		{"name": "red socks"},
		{"name": "green socks"},
		{"name": "blue socks"}
	]
}`),
		},
		{
			name: "failed",
			setup: func() {
				uc.On("List", testutils.OfTypeContext).Return(nil, errors.New("db error")).Once()
			},
			statusCode: http.StatusInternalServerError,
			res:        testutils.InternalServerErrorBytes,
		},
	}

	for _, tc := range tests {
		tc := tc
		tc.setup()

		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			handler := v1.NewProductHandler(uc)
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/v1/products", nil)

			handler.List(w, r)
			res := w.Result()
			defer res.Body.Close()

			assert.Equal(tc.statusCode, res.StatusCode, "status code does not match")
			got, err := ioutil.ReadAll(res.Body)
			assert.Nil(err)
			testutils.CmpJSON(t, tc.res, got)
		})
	}
}
