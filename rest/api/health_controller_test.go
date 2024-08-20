package api_test

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alextanhongpin/go-api-test/config"
	"github.com/alextanhongpin/go-api-test/rest/api"
	"github.com/alextanhongpin/testdump/httpdump"
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

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/health", nil)
	h := httpdump.HandlerFunc(t, handler, httpdump.IgnoreResponseFields("uptime", "startAt", "buildAt"))
	h.ServeHTTP(w, r)
}
