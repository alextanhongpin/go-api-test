package api_test

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alextanhongpin/core/test/testutil"
	"github.com/alextanhongpin/go-api-test/config"
	"github.com/alextanhongpin/go-api-test/rest/api"
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

	r := httptest.NewRequest("GET", "/health", nil)
	testutil.DumpHTTP(t, r, handler, testutil.IgnoreFields("uptime", "startAt", "buildAt"))
}
