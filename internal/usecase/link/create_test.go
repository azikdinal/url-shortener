package link

import (
	"context"
	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/url"
	"shorten/internal/usecase/link/mocks"
	domain "shorten/internal/domain/link"
	"testing"
)

func mustParseURL(raw string) *url.URL {
	u, err := url.Parse(raw)
	if err != nil {
		panic(err)
	}
	return u
}

func TestLinkUseCase_Create_Success(t *testing.T) {
	ctx := context.Background()

	storage := new(mocks.LinkStorage)
	storage.
		On("SaveFullURL", mock.Anything, "https://example.com").
		Return(int64(42), nil).
		Once()

	uc := &LinkUseCase{
		linkStorage: storage,
		publicHost:  mustParseURL("https://short.ly/"),
	}

	result, err := uc.Create(ctx, "https://example.com")

	require.NoError(t, err)

	expected := "https://short.ly/" + domain.GenerateShortCode(42)
	require.Equal(t, expected, result)

	storage.AssertExpectations(t)
}
