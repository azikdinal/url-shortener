package config

import "time"

type PostgresConfig struct {
	Host     string `envconfig:"PG_HOST"               default:"localhost"`
	Port     int    `envconfig:"PG_PORT"               default:"5432"`
	User     string `envconfig:"PG_USER"               default:"postgres"`
	Password string `envconfig:"PG_PASSWORD"           default:"postgres"`
	DBName   string `envconfig:"PG_DBNAME"             default:"postgres"`
	SSLMode  string `envconfig:"PG_SSL_MODE"           default:"disable"`

	MaxOpenConn     int           `envconfig:"PG_MAX_OPEN_CONN"      default:"25"`
	MaxIdleConn     int           `envconfig:"PG_MAX_IDLE_CONN"      default:"25"`
	ConnMaxLifetime time.Duration `envconfig:"PG_CONN_MAX_LIFETIME"  default:"5m"`
}
