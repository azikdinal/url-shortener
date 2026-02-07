package link

import (
	"context"
	domain "shorten/internal/domain/link"
)

func (uc *LinkUseCase) Create(
	ctx context.Context,
	fullURL string,
) (string, error) {

	link, err := domain.NewLink(fullURL)
	if err != nil {
		return "", err
	}

	err = uc.linkStorage.Save(ctx, link.ShortCode, link.FullURL)
	if err != nil {
		return "", err
	}

	return link.ShortCode, nil
}
