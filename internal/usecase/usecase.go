package usecase

import (
	"context"
	"shorten/internal/domain"
	"shorten/internal/repo"
)

type Usecase interface {
	GetByCode(ctx context.Context, sc domain.ShortCode) (domain.FullURL, error)
	Create(ctx context.Context, fu domain.FullURL) (domain.ShortCode, error)
}

type usecase struct {
	repo repo.Repo
}

func New(r repo.Repo) Usecase {
	return &usecase{
		repo: r,
	}
}

func (uc *usecase) GetByCode(
	ctx context.Context,
	sc domain.ShortCode,
) (domain.FullURL, error) {

	fu, err := uc.repo.GetByCode(ctx, sc)
	if err != nil {
		return "", err
	}

	return fu, err
}

func (uc *usecase) Create(
	ctx context.Context,
	fu domain.FullURL,
) (domain.ShortCode, error) {

	link := domain.NewLinkWithGeneratedCode(fu)

	err := uc.repo.Save(ctx, link)
	if err != nil {
		return "", err
	}

	return link.ShortCode(), nil
}
