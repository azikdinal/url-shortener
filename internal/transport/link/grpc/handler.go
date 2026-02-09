package grpc

import (
	"context"
	pb "shorten/gen/grpc/v1"
)

type LinkUseCase interface {
	GetFullURL(ctx context.Context, shortURL string) (string, error)
	Create(ctx context.Context, fullURL string) (string, error)
}

type LinkHandler struct {
	pb.UnimplementedShortenServiceServer
	linkUseCase LinkUseCase
}

func New(linkUseCase LinkUseCase) *LinkHandler {
	return &LinkHandler{
		linkUseCase: linkUseCase,
	}
}
