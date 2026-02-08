package link

import (
	"context"
	"errors"
	"net/url"
	"path"
	domain "shorten/internal/domain/link"
)

func (uc *LinkUseCase) GetFullURL(
	ctx context.Context,
	shortURL string,
) (string, error) {

	// Проверка валидности URL
	u, err := url.Parse(shortURL)
	if err != nil {
		return "", errIncorrectURL
	}
	// Проверка указанного хоста
	parsedHost := u.Scheme + "://" + u.Host + "/"
	if parsedHost != uc.publicHost.String() {
		return "", errIncorrectHost
	}

	shortCode := path.Base(u.Path)

	id := domain.ParseShortCode(shortCode)

	fullURL, err := uc.linkStorage.GetFullURL(ctx, id)
	if err != nil {
		return "", err
	}

	return fullURL, nil
}

var errIncorrectURL = errors.New("Неправильный формат URL")
var errIncorrectHost = errors.New("Неправильный хост")
