package app

import (
	"net"

	pb "shorten/gen/shorten/v1"
	"shorten/internal/config"
	"shorten/internal/router"
	grpcHandler "shorten/internal/transport/grpc"
	"shorten/internal/transport/rest"
	"shorten/internal/usecase"

	"google.golang.org/grpc"
)

func BuildHTTP(uc usecase.Usecase, cfg *config.Config) *router.Router {
	h := rest.NewHandler(uc, cfg.PublicHost)
	return router.New(h)
}

func BuildGRPC(
	uc usecase.Usecase,
	cfg *config.Config,
) (*grpc.Server, net.Listener, error) {

	lis, err := net.Listen("tcp", cfg.GRPC.Address)
	if err != nil {
		return nil, nil, err
	}

	s := grpc.NewServer()
	pb.RegisterShortenServiceServer(
		s,
		grpcHandler.NewHandler(uc, cfg.PublicHost),
	)
	return s, lis, nil
}
