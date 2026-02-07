package link

import (
	"context"
)

type LinkStorage interface {
	GetFullURL(ctx context.Context, fullURL string) (string, error)
	Save(ctx context.Context, shortCode string, fullURL string) error
}

type LinkUseCase struct {
	linkStorage LinkStorage
}

func New(s LinkStorage) *LinkUseCase {
	return &LinkUseCase{
		linkStorage: s,
	}
}
