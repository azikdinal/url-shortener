package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"shorten/internal/domain"
)

type Usecase struct {
	mock.Mock
}

func (m *Usecase) GetByCode(ctx context.Context, sc domain.ShortCode) (domain.FullURL, error) {
	args := m.Called(ctx, sc)
	return args.Get(0).(domain.FullURL), args.Error(1)
}

func (m *Usecase) Create(ctx context.Context, fu domain.FullURL) (domain.ShortCode, error) {
	args := m.Called(ctx, fu)
	return args.Get(0).(domain.ShortCode), args.Error(1)
}
