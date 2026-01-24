package domain

import (
	"errors"
	"net/url"
)

type FullURL string

func NewFullURL(raw string) (FullURL, error) {
	u, err := url.ParseRequestURI(raw)
	if err != nil {
		return "", err
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return "", errors.New("Неправильный протокол ссылки")
	}

	if u.Host == "" {
		return "", errors.New("У ссылки нет хоста")
	}

	return FullURL(u.String()), nil
}
