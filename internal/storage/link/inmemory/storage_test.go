package inmemory_test

// import (
// 	"context"
// 	domain "shorten/internal/domain/link"
// 	storage "shorten/internal/storage/link/inmemory"
// 	"testing"
//
// 	"github.com/stretchr/testify/require"
// )
//
// func TestMemoryRepo_SaveAndGet(t *testing.T) {
// 	r := storage.New()
//
// 	sc := "sdfmv_dsfd"
// 	fu := "https://example.com"
// 	link := domain.NewLink(fu)
// 	require.NoError(t, r.Save(context.Background(), link))
//
// 	got, err := r.GetByCode(context.Background(), sc)
// 	require.NoError(t, err)
// 	require.Equal(t, fu, got)
// }
