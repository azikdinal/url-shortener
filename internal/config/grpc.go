package config

type GRPCConfig struct {
	Address string `envconfig:"GRPC_URL" default:"localhost:8001"`
}
