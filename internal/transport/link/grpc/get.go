package grpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "shorten/gen/shorten/v1"
)

func (h *LinkHandler) Get(
	ctx context.Context,
	req *pb.GetByCodeRequest,
) (*pb.GetByCodeResponse, error) {

	shortCode := req.ShortCode

	fullURL, err := h.linkUseCase.GetFullURL(ctx, shortCode)
	if err != nil {
		return nil, status.Error(codes.NotFound, "link not found")
	}

	return &pb.GetByCodeResponse{
		FullUrl: string(fullURL),
	}, nil
}
