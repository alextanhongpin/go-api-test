package api_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alextanhongpin/go-api-test/config"
	"github.com/alextanhongpin/go-api-test/internal/testutils"
	"github.com/alextanhongpin/go-api-test/rest/api"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestHealthController(t *testing.T) {
	now := time.Now()
	handler := api.NewHealthController(&config.Config{
		Name:    "test",
		Version: "0.0.1",
		BuildAt: now,
		StartAt: now,
		VCSRef:  "xyz",
		VCSURL:  "http://xyz",
	})

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/health", nil)

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

	testutils.CmpJSON(t, want, b, cmpopts.IgnoreMapEntries(func(k string, v any) bool {
		// Skip fields that starts with "uptime", because the result is
		// unpredictable.
		return k == "uptime"
	}))
}
