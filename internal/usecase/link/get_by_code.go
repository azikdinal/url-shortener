package link

import (
	"context"
)

func (uc *LinkUseCase) GetFullURL(
	ctx context.Context,
	shortURL string,
) (string, error) {

	fullURL, err := uc.linkStorage.GetFullURL(ctx, shortURL)
	if err != nil {
		return "", err
	}

	return fullURL, err
}
