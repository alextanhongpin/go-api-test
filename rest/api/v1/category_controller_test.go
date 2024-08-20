package v1_test

import (
	"context"
	"net/http/httptest"
	"testing"

	v1 "github.com/alextanhongpin/go-api-test/rest/api/v1"
	"github.com/alextanhongpin/go-api-test/rest/contextkey"
	"github.com/alextanhongpin/testdump/httpdump"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func TestCategoryControllerCreate(t *testing.T) {
	ctx := context.Background()
	ctx = contextkey.SetUserID(ctx, uuid.New())

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/v1/categories", nil)
	r = r.WithContext(ctx)
	handler := new(v1.CategoryController).Create
	h := httpdump.HandlerFunc(t, handler)
	h.ServeHTTP(w, r)
}

func TestCategoryControllerShow(t *testing.T) {

	// We need to inject the URL params manually when using chi router.
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/v1/categories/1", nil)
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

	handler := new(v1.CategoryController).Show
	h := httpdump.HandlerFunc(t, handler)
	h.ServeHTTP(w, r)
}

func TestCategoryControllerList(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/v1/categories", nil)
	handler := new(v1.CategoryController).List
	h := httpdump.HandlerFunc(t, handler)
	h.ServeHTTP(w, r)
}

func TestCategoryControllerUpdate(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("PATCH", "/v1/categories", nil)
	handler := new(v1.CategoryController).Update
	h := httpdump.HandlerFunc(t, handler)
	h.ServeHTTP(w, r)
}

func TestCategoryControllerDelete(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/v1/categories", nil)
	handler := new(v1.CategoryController).Delete
	h := httpdump.HandlerFunc(t, handler)
	h.ServeHTTP(w, r)
}
