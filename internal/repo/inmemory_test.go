package repo

import (
	"context"
	"shorten/internal/domain"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMemoryRepo_SaveAndGet(t *testing.T) {
	r := NewMemoryRepo()

	link := domain.NewLink(domain.ShortCode("sdfmv_dsfd"), domain.FullURL("https://example.com"))
	require.NoError(t, r.Save(context.Background(), link))

	got, err := r.GetByCode(context.Background(), link.ShortCode())
	require.NoError(t, err)
	require.Equal(t, link, got)
}
