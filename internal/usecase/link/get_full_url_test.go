package link

import (
	"context"
	mock "github.com/stretchr/testify/mock"
	require "github.com/stretchr/testify/require"
	domain "shorten/internal/domain/link"
	"shorten/internal/usecase/link/mocks"
	"testing"
)

func TestLinkUseCase_GetFullURL_Success(t *testing.T) {
	ctx := context.Background()

	id := int64(42)
	code := domain.GenerateShortCode(id)
	shortURL := "https://short.ly/" + code
	expectedFullURL := "https://example.com/page"

	storage := new(mocks.LinkStorage)
	storage.
		On("GetFullURL", mock.Anything, id).
		Return(expectedFullURL, nil).
		Once()

	uc := &LinkUseCase{
		linkStorage: storage,
		publicHost:  mustParseURL("https://short.ly/"),
	}

	result, err := uc.GetFullURL(ctx, shortURL)

	require.NoError(t, err)
	require.Equal(t, expectedFullURL, result)

	storage.AssertExpectations(t)
}

func TestLinkUseCase_GetFullURL_InvalidURL(t *testing.T) {
	uc := &LinkUseCase{
		publicHost: mustParseURL("https://short.ly/"),
	}

	result, err := uc.GetFullURL(context.Background(), "://bad-url")

	require.Error(t, err)
	require.ErrorIs(t, err, errIncorrectURL)
	require.Empty(t, result)
}

func TestLinkUseCase_GetFullURL_InvalidHost(t *testing.T) {
	id := int64(1)
	code := domain.GenerateShortCode(id)

	uc := &LinkUseCase{
		publicHost: mustParseURL("https://short.ly/"),
	}

	shortURL := "https://evil.com/" + code

	result, err := uc.GetFullURL(context.Background(), shortURL)

	require.Error(t, err)
	require.ErrorIs(t, err, errIncorrectHost)
	require.Empty(t, result)
}
