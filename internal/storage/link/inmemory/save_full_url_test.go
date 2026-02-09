package inmemory

import (
	"context"
	"testing"
)

func TestInmemoryStorage_SaveFullURL(t *testing.T) {
	s := &InmemoryStorage{}

	id, err := s.SaveFullURL(context.Background(), "https://example.com")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if id <= 0 {
		t.Fatalf("expected id > 0, got %d", id)
	}

	val, ok := s.data.Load(id)
	if !ok {
		t.Fatalf("value not stored")
	}

	if val.(string) != "https://example.com" {
		t.Fatalf("unexpected value: %v", val)
	}
}
