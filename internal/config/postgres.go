package config

import "time"

type PostgresConfig struct {
	Host     string `envconfig:"PG_HOST"`
	Port     int    `envconfig:"PG_PORT"`
	User     string `envconfig:"PG_USER"`
	Password string `envconfig:"PG_PASSWORD"`
	DBName   string `envconfig:"PG_DBNAME"`
	SSLMode  string `envconfig:"PG_SSL_MODE" default:"disable"`

	MaxOpenConn int `envconfig:"PG_MAX_OPEN_CONN" default:"10"`
	MinOpenConn int `envconfig:"PG_MIN_OPEN_CONN" default:"5"`

	MaxConnLifetime time.Duration `envconfig:"PG_MAX_LIFETIME" default:"5m"`
	MaxConnIdletime time.Duration `envconfig:"PG_MAX_IDLETIME" default:"5m"`
}
