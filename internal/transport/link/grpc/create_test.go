package grpc

import (
	"context"
	pb "shorten/gen/grpc/v1"
	"shorten/internal/transport/link/mocks"
	"testing"

	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestLinkHandler_Create_Success(t *testing.T) {
	ctx := context.Background()

	useCase := new(mocks.LinkUseCase)
	useCase.
		On("Create", mock.Anything, "https://example.com").
		Return("https://short.ly/abc123", nil).
		Once()

	handler := &LinkHandler{
		linkUseCase: useCase,
	}

	req := &pb.CreateRequest{
		FullUrl: "https://example.com",
	}

	resp, err := handler.Create(ctx, req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, "https://short.ly/abc123", resp.ShortUrl)

	useCase.AssertExpectations(t)
}
