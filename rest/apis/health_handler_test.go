package apis_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alextanhongpin/go-api-test/rest/apis"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestHealthHandler(t *testing.T) {
	now := time.Now()
	handler := apis.NewHealthHandler(&apis.HealthHandlerConfig{
		Name:    "test",
		Version: "0.0.1",
		BuildAt: now,
		StartAt: now,
		VCSRef:  "xyz",
		VCSURL:  "http://xyz",
	})

	r, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()

	handler.Show(w, r)
	res := w.Result()
	defer res.Body.Close()

	if want, got := http.StatusOK, res.StatusCode; want != got {
		t.Fatalf("status code: want %d, got %d", want, got)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	tb, err := json.Marshal(now)
	if err != nil {
		t.Error(err)
	}

	want := []byte(fmt.Sprintf(`{
	"serviceName": "test",
	"version": "0.0.1",
	"buildAt": %s,
	"startAt": %s,
	"vcsRef": "xyz",
	"vcsUrl": "http://xyz"
}`,
		tb,
		tb,
	))

	cmpJSON(t, want, b, cmpopts.IgnoreMapEntries(func(k string, v any) bool {
		// Skip fields that starts with "uptime", because the result is
		// unpredictable.
		return k == "uptime"
	}))
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
