package link

import (
	"fmt"

	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "shorten/gen/shorten/v1"
)

func (h *LinkHandler) Create(
	ctx context.Context,
	req *pb.CreateRequest,
) (*pb.CreateResponse, error) {

	fullURL := req.FullUrl

	shortURL, err := h.linkUseCase.Create(ctx, fullURL)
	if err != nil {
		return nil,
			status.Error(
				codes.NotFound,
				fmt.Sprintf("link fetching is failed: %s", err.Error()),
			)
	}

	return &pb.CreateResponse{
		ShortUrl: shortURL,
	}, nil
}
