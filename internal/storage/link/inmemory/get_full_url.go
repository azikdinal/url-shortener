package inmemory

import (
	"context"
	"errors"
)

func (s *InMemoryStorage) GetFullURL(
	ctx context.Context,
	shortCode string,
) (string, error) {

	s.mu.RLock()
	defer s.mu.RUnlock()

	fullURL, ok := s.data[shortCode]
	if !ok {
		return "", errors.New("link not found in memory")
	}

	return fullURL, nil
}
