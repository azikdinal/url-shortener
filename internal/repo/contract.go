package repo

import (
	"context"
	"shorten/internal/domain"
)

type Repo interface {
	GetByCode(ctx context.Context, sc domain.ShortCode) (domain.FullURL, error)
	Save(ctx context.Context, link *domain.Link) error
}
