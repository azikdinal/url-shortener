package config

import "time"

type HTTPConfig struct {
	Address      string        `envconfig:"HTTP_URL"           default:"localhost:8000"`
	ReadTimeout  time.Duration `envconfig:"HTTP_READ_TIMEOUT"  default:"5s"`
	WriteTimeout time.Duration `envconfig:"HTTP_WRITE_TIMEOUT" default:"5s"`
}
