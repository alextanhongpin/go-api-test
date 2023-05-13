package api_test

import (
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
	}).Show

	opt := cmpopts.IgnoreMapEntries(func(k string, v any) bool {
		// Skip fields that starts with "uptime", because the result is
		// unpredictable.
		switch k {
		case "uptime":
			return true
		case "startAt", "buildAt":
			return testutils.IsJsonTime(t, v)
		default:
			return false
		}
	})

	r := httptest.NewRequest("GET", "/health", nil)
	testutils.HTTPSnapshot(t, r, handler, "./testdata/get_health_response.json", http.StatusOK, opt)
}
