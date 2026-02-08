package inmemory

import (
	"context"
	"sync/atomic"
)

// Сохранениние fullURL в памяти для последующего update
func (s *InmemoryStorage) SaveFullURL(ctx context.Context, fullURL string) (int64, error) {
	id := atomic.AddInt64(&s.nextID, 1)

	s.data.Store(id, fullURL)
	return id, nil
}
