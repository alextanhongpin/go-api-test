package config

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Name    string    `envconfig:"APP_NAME" required:"true" desc:"The app name"`
	Version string    `envconfig:"APP_VERSION" required:"true" desc:"The app version"`
	BuildAt time.Time `envconfig:"BUILD_DATE" required:"true" desc:"The build date of the container"`
	StartAt time.Time `ignored:"true"`
	VCSRef  string    `envconfig:"VCS_REF" required:"true" desc:"The version source control commit hash"`
	VCSURL  string    `envconfig:"VCS_URL" required:"true" desc:"The version source control URL"`
	JWT     JWTConfig
}

func New() *Config {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err)
	}

	// Custom overrides.
	cfg.StartAt = time.Now() // The time the server starts.

	return &cfg
}

func (c *Config) Uptime() time.Duration {
	return time.Since(c.StartAt)
}
