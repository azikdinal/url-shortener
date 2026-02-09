package app

import (
	"net"

	pb "shorten/gen/grpc/v1"
	"shorten/internal/config"
	grpcHandler "shorten/internal/transport/link/grpc"

	"google.golang.org/grpc"
)

func buildGRPCServer(
	uc grpcHandler.LinkUseCase,
	cfg *config.Config,
) (*grpc.Server, net.Listener, error) {

	lis, err := net.Listen("tcp", cfg.GRPC_URL)
	if err != nil {
		return nil, nil, err
	}

	s := grpc.NewServer()
	pb.RegisterShortenServiceServer(
		s,
		grpcHandler.New(uc),
	)
	return s, lis, nil
}
