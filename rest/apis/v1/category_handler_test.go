package v1_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	v1 "github.com/alextanhongpin/go-api-test/rest/apis/v1"
	"github.com/google/go-cmp/cmp"
)

func TestCategoryHandlerShow(t *testing.T) {
	r, err := http.NewRequest("GET", "/v1/categories", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	handler := new(v1.CategoryHandler).Show
	handler(w, r)

	res := w.Result()
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	want := []byte(`{
	"data": {
		"name": "Toys"
	}
}`)

	cmpJSON(t, want, b)
}

func cmpJSON(t *testing.T, a, b []byte, opts ...cmp.Option) {
	var l, r map[string]any
	if err := json.Unmarshal(a, &l); err != nil {
		t.Error(err)
	}
	if err := json.Unmarshal(b, &r); err != nil {
		t.Error(err)
	}

	if diff := cmp.Diff(l, r, opts...); diff != "" {
		t.Errorf("want(+), got(-): %s", diff)
	}
}
