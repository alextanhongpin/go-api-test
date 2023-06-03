package v1_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/alextanhongpin/core/test/testutil"
	v1 "github.com/alextanhongpin/go-api-test/rest/api/v1"
	"github.com/alextanhongpin/go-api-test/rest/contextkey"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func TestCategoryControllerCreate(t *testing.T) {
	ctx := context.Background()
	ctx = contextkey.SetUserID(ctx, uuid.New())

	r := httptest.NewRequest("POST", "/v1/categories", nil)
	r = r.WithContext(ctx)
	handler := new(v1.CategoryController).Create
	testutil.DumpHTTP(t, r, handler)
}

func TestCategoryControllerShow(t *testing.T) {

	// We need to inject the URL params manually when using chi router.
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")

	r := httptest.NewRequest("GET", "/v1/categories/1", nil)
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

	handler := new(v1.CategoryController).Show
	testutil.DumpHTTP(t, r, handler)
}

func TestCategoryControllerList(t *testing.T) {
	r := httptest.NewRequest("GET", "/v1/categories", nil)
	handler := new(v1.CategoryController).List
	testutil.DumpHTTP(t, r, handler)
}

func TestCategoryControllerUpdate(t *testing.T) {
	r := httptest.NewRequest("PATCH", "/v1/categories", nil)
	handler := new(v1.CategoryController).Update
	testutil.DumpHTTP(t, r, handler)
}

func TestCategoryControllerDelete(t *testing.T) {
	r := httptest.NewRequest("DELETE", "/v1/categories", nil)
	handler := new(v1.CategoryController).Delete
	testutil.DumpHTTP(t, r, handler)
}
