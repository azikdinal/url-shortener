package link

import (
	"context"
	domain "shorten/internal/domain/link"
)

func (uc *LinkUseCase) Create(
	ctx context.Context,
	fullURL string,
) (string, error) {

	// Сохраняем запись для получения id
	id, err := uc.linkStorage.SaveFullURL(ctx, fullURL)
	if err != nil {
		return "", err
	}

	// Генерируем shortCode по ID
	shortCode := domain.GenerateShortCode(id)

	return uc.publicHost.String() + shortCode, nil
}
