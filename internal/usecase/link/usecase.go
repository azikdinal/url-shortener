package link

import (
	"context"
)

type LinkStorage interface {
	GetFullURL(ctx context.Context, fullURL string) (string, error)
	Save(ctx context.Context, fullURL string) (int64, error)
	Update(ctx context.Context, id int64, shortCode string) error
}

type ShortCodeGenerator interface {
	Generate(id int64) (string, error)
}

type LinkUseCase struct {
	linkStorage LinkStorage
	scGen       ShortCodeGenerator
}

func New(s LinkStorage, scGen ShortCodeGenerator) *LinkUseCase {
	return &LinkUseCase{
		linkStorage: s,
		scGen:       scGen,
	}
}
