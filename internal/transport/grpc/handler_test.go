package grpc_test

import (
	"context"
	"testing"

	pb "shorten/gen/shorten/v1"
	"shorten/internal/domain"
	grpcHandler "shorten/internal/transport/grpc"
	"shorten/internal/usecase/mocks"

	"github.com/stretchr/testify/require"
)

func TestGRPCHandler_Create(t *testing.T) {
	host := "http://sh.by"
	shortCode := "fjmi_dsf87"
	sc := domain.ShortCode(shortCode)
	fullURL := "https://example.com"
	fu := domain.FullURL(fullURL)

	mockUC := new(mocks.Usecase)
	handler := grpcHandler.NewHandler(mockUC, host)

	ctx := context.Background()
	req := &pb.CreateRequest{
		FullUrl: fullURL,
	}

	mockUC.On("Create", ctx, fu).Return(sc, nil)

	resp, err := handler.Create(ctx, req)
	require.NoError(t, err)
	require.Equal(t, host+"/"+shortCode, resp.ShortUrl)
}

func TestGRPCHandler_GetByCode(t *testing.T) {
	host := "http://sh.by"
	fullURL := "https://example.com/"
	shortCode := "advmi_d4_d"
	sc := domain.ShortCode(shortCode)
	fu, _ := domain.NewFullURL(fullURL)

	mockUC := new(mocks.Usecase)
	handler := grpcHandler.NewHandler(mockUC, host)

	ctx := context.Background()
	req := &pb.GetByCodeRequest{
		ShortCode: shortCode,
	}

	mockUC.On("GetByCode", ctx, sc).Return(fu, nil)

	resp, err := handler.GetByCode(ctx, req)
	require.NoError(t, err)
	require.Equal(t, fullURL, resp.FullUrl)
}
