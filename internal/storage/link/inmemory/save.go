package inmemory

import (
	"context"
)

func (r *InMemoryStorage) Save(ctx context.Context, shortCode string, fullURL string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	key := shortCode
	r.data[key] = fullURL

	return nil
}
