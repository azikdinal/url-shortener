package repo

import (
	"context"
	"errors"
	"shorten/internal/domain"
	"sync"
)

type MemoryRepo struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewMemoryRepo() Repo {
	return &MemoryRepo{
		mu:   sync.RWMutex{},
		data: map[string]string{},
	}
}

func (r *MemoryRepo) GetByCode(
	ctx context.Context,
	sc domain.ShortCode,
) (domain.FullURL, error) {

	r.mu.RLock()
	defer r.mu.RUnlock()

	fullURLStr, ok := r.data[string(sc)]
	if !ok {
		return "", errors.New("link not found in memory")
	}

	fu, err := domain.NewFullURL(fullURLStr)
	if err != nil {
		return "", errors.New("not valid value of fullURL returned from memory")
	}

	return fu, nil
}

func (r *MemoryRepo) Save(ctx context.Context, link *domain.Link) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	key := string(link.ShortCode())
	r.data[key] = string(link.FullURL())

	return nil
}
