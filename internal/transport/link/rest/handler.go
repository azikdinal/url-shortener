package rest

import (
	"context"
)

type LinkUseCase interface {
	GetFullURL(ctx context.Context, shortCode string) (string, error)
	Create(ctx context.Context, fullURL string) (string, error)
}

type LinkHandler struct {
	linkUseCase LinkUseCase
	publicHost  string
}

func New(uc LinkUseCase, ph string) *LinkHandler {
	return &LinkHandler{
		linkUseCase: uc,
		publicHost:  ph,
	}
}
