package inmemory

import (
	"context"
	"errors"
)

func (r *InMemoryStorage) GetFullURL(
	ctx context.Context,
	shortCode string,
) (string, error) {

	r.mu.RLock()
	defer r.mu.RUnlock()

	fullURL, ok := r.data[shortCode]
	if !ok {
		return "", errors.New("link not found in memory")
	}

	return fullURL, nil
}
