package apis

import (
	"net/http"

	"github.com/alextanhongpin/go-api-test/config"
	"github.com/alextanhongpin/go-core-microservice/http/encoding"
)

type HealthHandler struct {
	cfg *config.Config
}

func NewHealthHandler(cfg *config.Config) *HealthHandler {
	return &HealthHandler{
		cfg: cfg,
	}
}

func (c *HealthHandler) Show(w http.ResponseWriter, r *http.Request) {
	res := Health{
		BuildAt:     c.cfg.BuildAt,
		StartAt:     c.cfg.StartAt,
		Uptime:      c.cfg.Uptime().String(),
		VCSRef:      c.cfg.VCSRef,
		VCSURL:      c.cfg.VCSURL,
		Version:     c.cfg.Version,
		ServiceName: c.cfg.Name,
		// You can also add service health, such as db, redis, external services
		// etc. It can be a simple ping, and a string message indicating the health
		// of the service.
	}

	encoding.EncodeJSON(w, res, http.StatusOK)
}
