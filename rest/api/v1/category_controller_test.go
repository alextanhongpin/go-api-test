package v1_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alextanhongpin/go-api-test/internal/testutils"
	v1 "github.com/alextanhongpin/go-api-test/rest/api/v1"
	"github.com/alextanhongpin/go-api-test/rest/security"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func TestCategoryControllerCreate(t *testing.T) {
	ctx := context.Background()
	ctx = security.ContextWithClaims(ctx, jwt.MapClaims{
		"exp": time.Now().Add(1 * time.Hour).Unix(),
		"sub": uuid.New().String(),
	})

	r := httptest.NewRequest("POST", "/v1/categories", nil)
	r = r.WithContext(ctx)
	handler := new(v1.CategoryController).Create
	testutils.HTTPSnapshot(t, r, handler, "./testdata/create_category_response.json", http.StatusInternalServerError)
}

func TestCategoryControllerShow(t *testing.T) {

	// We need to inject the URL params manually when using chi router.
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")

	r := httptest.NewRequest("GET", "/v1/categories/1", nil)
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

	handler := new(v1.CategoryController).Show
	testutils.HTTPSnapshot(t, r, handler, "./testdata/show_category_response.json", http.StatusInternalServerError)
}

func TestCategoryControllerList(t *testing.T) {
	r := httptest.NewRequest("GET", "/v1/categories", nil)
	handler := new(v1.CategoryController).List
	testutils.HTTPSnapshot(t, r, handler, "./testdata/list_category_response.json", http.StatusInternalServerError)
}

func TestCategoryControllerUpdate(t *testing.T) {
	r := httptest.NewRequest("PATCH", "/v1/categories", nil)
	handler := new(v1.CategoryController).Update
	testutils.HTTPSnapshot(t, r, handler, "./testdata/update_category_response.json", http.StatusInternalServerError)
}

func TestCategoryControllerDelete(t *testing.T) {
	r := httptest.NewRequest("DELETE", "/v1/categories", nil)
	handler := new(v1.CategoryController).Delete
	testutils.HTTPSnapshot(t, r, handler, "./testdata/delete_category_response.json", http.StatusInternalServerError)
}
