package grpc

import (
	"context"
	pb "shorten/gen/shorten/v1"
	"shorten/internal/transport/link/mocks"
	"testing"

	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestLinkHandler_Get_Success(t *testing.T) {
	ctx := context.Background()

	useCase := new(mocks.LinkUseCase)
	useCase.
		On("GetFullURL", mock.Anything, "abc123").
		Return("https://example.com", nil).
		Once()

	handler := &LinkHandler{
		linkUseCase: useCase,
	}

	req := &pb.GetByCodeRequest{
		ShortCode: "abc123",
	}

	resp, err := handler.Get(ctx, req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, "https://example.com", resp.FullUrl)

	useCase.AssertExpectations(t)
}
