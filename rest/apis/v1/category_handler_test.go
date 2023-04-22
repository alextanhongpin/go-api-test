package v1_test

import (
	"context"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	tu "github.com/alextanhongpin/go-api-test/internal/testutils"
	v1 "github.com/alextanhongpin/go-api-test/rest/apis/v1"
	"github.com/go-chi/chi/v5"
)

func TestCategoryHandlerCreate(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/v1/categories", nil)
	handler := new(v1.CategoryHandler).Create
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

func TestCategoryHandlerShow(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/v1/categories/1", nil)

	// We need to inject the URL params manually when using chi router.
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
	handler := new(v1.CategoryHandler).Show
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

func TestCategoryHandlerList(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/v1/categories", nil)
	handler := new(v1.CategoryHandler).List
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

func TestCategoryHandlerUpdate(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("PATCH", "/v1/categories", nil)
	handler := new(v1.CategoryHandler).Update
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

func TestCategoryHandlerDelete(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/v1/categories", nil)
	handler := new(v1.CategoryHandler).Delete
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
