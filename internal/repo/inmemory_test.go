package repo

import (
	"context"
	"shorten/internal/domain"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMemoryRepo_SaveAndGet(t *testing.T) {
	r := NewMemoryRepo()

	sc := domain.ShortCode("sdfmv_dsfd")
	fu := domain.FullURL("https://example.com")
	link := domain.NewLink(sc, fu)
	require.NoError(t, r.Save(context.Background(), link))

	got, err := r.GetByCode(context.Background(), sc)
	require.NoError(t, err)
	require.Equal(t, fu, got)
}
