package link

import (
	"context"
	"net/url"
)

type LinkStorage interface {
	GetFullURL(ctx context.Context, id int64) (string, error)
	SaveFullURL(ctx context.Context, fullURL string) (int64, error)
}

type LinkUseCase struct {
	linkStorage LinkStorage
	publicHost  *url.URL
}

func New(s LinkStorage, ph *url.URL) *LinkUseCase {
	return &LinkUseCase{
		linkStorage: s,
		publicHost:  ph,
	}
}
