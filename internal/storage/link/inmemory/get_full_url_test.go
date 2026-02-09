package inmemory

import (
	"context"
	"errors"
	"testing"
)

func TestInmemoryStorage_GetFullURL(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name      string
		setup     func(s *InmemoryStorage)
		id        int64
		wantURL   string
		wantError error
	}{
		{
			name: "success",
			setup: func(s *InmemoryStorage) {
				s.data.Store(int64(1), "https://example.com")
			},
			id:        1,
			wantURL:   "https://example.com",
			wantError: nil,
		},
		{
			name: "id not found",
			setup: func(s *InmemoryStorage) {
			},
			id:        42,
			wantURL:   "",
			wantError: errNotFound,
		},
		{
			name: "invalid value type",
			setup: func(s *InmemoryStorage) {
				s.data.Store(int64(2), 12345) // не string
			},
			id:        2,
			wantURL:   "",
			wantError: errNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := &InmemoryStorage{}
			tt.setup(storage)

			url, err := storage.GetFullURL(ctx, tt.id)

			if !errors.Is(err, tt.wantError) {
				t.Fatalf("expected error %v, got %v", tt.wantError, err)
			}

			if url != tt.wantURL {
				t.Fatalf("expected url %q, got %q", tt.wantURL, url)
			}
		})
	}
}
