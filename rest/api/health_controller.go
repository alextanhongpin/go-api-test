package api

import (
	"net/http"

	"github.com/alextanhongpin/core/http/response"
	"github.com/alextanhongpin/go-api-test/config"
)

type HealthController struct {
	cfg *config.Config
}

func NewHealthController(cfg *config.Config) *HealthController {
	return &HealthController{
		cfg: cfg,
	}
}

func (c *HealthController) Show(w http.ResponseWriter, r *http.Request) {
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

	response.JSON(w, res, http.StatusOK)
}
