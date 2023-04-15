package v1_test

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/alextanhongpin/go-api-test/internal/testutils"
	v1 "github.com/alextanhongpin/go-api-test/rest/apis/v1"
)

func TestCategoryHandlerShow(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/v1/categories", nil)
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

	testutils.CmpJSON(t, want, got)
}
