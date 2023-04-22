package config

type JWTConfig struct {
	Secret string `envconfig:"JWT_SECRET" required:"true" desc:"The JWT secret"`
}
