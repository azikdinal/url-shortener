package inmemory

import (
	"context"
)

func (s *InMemoryStorage) Save(ctx context.Context, shortCode string, fullURL string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	key := shortCode
	s.data[key] = fullURL

	return nil
}
