package link

import (
	"context"
)

func (uc *LinkUseCase) Create(
	ctx context.Context,
	fullURL string,
) (string, error) {

	// Сохраняем запись для получения id
	id, err := uc.linkStorage.Save(ctx, fullURL)
	if err != nil {
		return "", err
	}

	// Генерируем shortCode по ID
	shortCode, err := uc.scGen.Generate(id)
	if err != nil {
		return "", err
	}

	// Добавляем shortCode к записи
	err = uc.linkStorage.Update(ctx, id, shortCode)
	if err != nil {
		return "", err
	}

	return shortCode, nil
}
