package config

import "time"

type HTTPConfig struct {
	Address      string        `envconfig:"BACKEND_URL"           default:"localhost:8000"`
	PublicHost   string        `envconfig:"PUBLIC_HOST"           default:"http://localhost:8000/"`
	ReadTimeout  time.Duration `envconfig:"BACKEND_READ_TIMEOUT"  default:"5s"`
	WriteTimeout time.Duration `envconfig:"BACKEND_WRITE_TIMEOUT" default:"5s"`
}
