package apis

import (
	"net/http"
	"time"

	"github.com/alextanhongpin/go-core-microservice/http/encoding"
)

type HealthHandler struct {
	cfg *HealthHandlerConfig
}

type HealthHandlerConfig struct {
	Name    string
	Version string
	BuildAt time.Time
	StartAt time.Time
	VCSRef  string
	VCSURL  string
}

func NewHealthHandler(cfg *HealthHandlerConfig) *HealthHandler {
	return &HealthHandler{
		cfg: cfg,
	}
}

func (c *HealthHandler) Show(w http.ResponseWriter, r *http.Request) {
	res := Health{
		BuildAt:     c.cfg.BuildAt,
		StartAt:     c.cfg.StartAt,
		Uptime:      time.Since(c.cfg.StartAt).String(),
		VCSRef:      c.cfg.VCSRef,
		VCSURL:      c.cfg.VCSURL,
		Version:     c.cfg.Version,
		ServiceName: c.cfg.Name,
	}

	encoding.EncodeJSON(w, res, http.StatusOK)
}
