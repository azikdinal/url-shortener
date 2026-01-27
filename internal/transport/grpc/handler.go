package grpc

import (
	"context"
	"fmt"
	pb "shorten/gen/shorten/v1"
	"shorten/internal/domain"
	"shorten/internal/usecase"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	pb.UnimplementedShortenServiceServer
	uc   usecase.Usecase
	host string
}

func NewHandler(uc usecase.Usecase, host string) *Handler {
	return &Handler{
		uc:   uc,
		host: host,
	}
}

func (h *Handler) GetByCode(
	ctx context.Context,
	req *pb.GetByCodeRequest,
) (*pb.GetByCodeResponse, error) {

	sc, err := domain.NewShortCode(req.ShortCode)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "given link is invalid")
	}

	fu, err := h.uc.GetByCode(ctx, sc)
	if err != nil {
		return nil, status.Error(codes.NotFound, "link not found")
	}

	return &pb.GetByCodeResponse{
		FullUrl: string(fu),
	}, nil
}

func (h *Handler) Create(
	ctx context.Context,
	req *pb.CreateRequest,
) (*pb.CreateResponse, error) {

	fu, err := domain.NewFullURL(req.FullUrl)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "given source link is invalid")
	}

	sc, err := h.uc.Create(ctx, fu)
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("link fetching is failed: %s", err.Error()))
	}

	return &pb.CreateResponse{
		ShortUrl: h.host + "/" + string(sc),
	}, nil
}
