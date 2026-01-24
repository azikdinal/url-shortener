package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	HTTP      HTTPConfig
	Postgres  PostgresConfig
	StoreType StoreType `envconfig:"STORE_TYPE" default:"IN_MEMORY"`
}

func Load() (*config, error) {
	_ = godotenv.Load()
	var cfg config

	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
