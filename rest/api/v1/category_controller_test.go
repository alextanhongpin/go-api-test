package v1_test

import (
	"context"
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"time"

	tu "github.com/alextanhongpin/go-api-test/internal/testutils"
	v1 "github.com/alextanhongpin/go-api-test/rest/apis/v1"
	"github.com/alextanhongpin/go-api-test/rest/security"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func TestCategoryControllerCreate(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/v1/categories", nil)
	ctx := context.Background()
	ctx = security.AuthContext.WithValue(ctx, jwt.MapClaims{
		"exp": time.Now().Add(1 * time.Hour).Unix(),
		"sub": uuid.New().String(),
	})
	r = r.WithContext(ctx)
	handler := new(v1.CategoryController).Create
	handler(w, r)

	res := w.Result()
	defer res.Body.Close()

	got, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	want := tu.InternalServerErrorBytes
	tu.CmpJSON(t, want, got)
}

func TestCategoryControllerShow(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/v1/categories/1", nil)

	// We need to inject the URL params manually when using chi router.
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
	handler := new(v1.CategoryController).Show
	handler(w, r)

	res := w.Result()
	defer res.Body.Close()

	got, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	want := []byte(`{
	"data": {
		"name": "Toys"
	}
}`)

	tu.CmpJSON(t, want, got)
}

func TestCategoryControllerList(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/v1/categories", nil)
	handler := new(v1.CategoryController).List
	handler(w, r)

	res := w.Result()
	defer res.Body.Close()

	got, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	want := tu.InternalServerErrorBytes
	tu.CmpJSON(t, want, got)
}

func TestCategoryControllerUpdate(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("PATCH", "/v1/categories", nil)
	handler := new(v1.CategoryController).Update
	handler(w, r)

	res := w.Result()
	defer res.Body.Close()

	got, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	want := tu.InternalServerErrorBytes
	tu.CmpJSON(t, want, got)
}

func TestCategoryControllerDelete(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/v1/categories", nil)
	handler := new(v1.CategoryController).Delete
	handler(w, r)

	res := w.Result()
	defer res.Body.Close()

	got, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	want := tu.InternalServerErrorBytes
	tu.CmpJSON(t, want, got)
}
