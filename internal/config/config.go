package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HTTP          HTTPConfig
	GRPC_URL      string `envconfig:"GRPC_URL" default:"localhost:8001"`
	Postgres      PostgresConfig
	StoreType     StoreType `envconfig:"STORE_TYPE"  default:"IN_MEMORY"`
	PublicHost    string    `envconfig:"PUBLIC_HOST" default:"http://localhost:8000/"`
	ShortCodeSeed string    `envconfig:"SHORT_CODE_SEED" default:"dmfijh98&K"`
}

func Load() (*Config, error) {
	_ = godotenv.Load()
	var cfg Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
