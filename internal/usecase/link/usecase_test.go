package link_test

// import (
// 	"context"
// 	"testing"
//
// 	"shorten/internal/domain"
// 	"shorten/internal/usecase"
//
// 	"github.com/stretchr/testify/mock"
// 	"github.com/stretchr/testify/require"
// )
//
// type MockRepo struct {
// 	mock.Mock
// }
//
// func (m *MockRepo) Save(ctx context.Context, link *domain.Link) error {
// 	args := m.Called(ctx, link)
// 	return args.Error(0)
// }
//
// func (m *MockRepo) GetByCode(
// 	ctx context.Context,
// 	sc domain.ShortCode,
// ) (domain.FullURL, error) {
//
// 	args := m.Called(ctx, sc)
// 	return args.Get(0).(domain.FullURL), args.Error(1)
// }
//
// func TestUsecase_Create(t *testing.T) {
// 	repo := new(MockRepo)
// 	uc := usecase.New(repo)
//
// 	link := domain.NewLink(domain.ShortCode("dsfj_d_did"), domain.FullURL("https://example.com"))
//
// 	repo.On("Save", mock.Anything, mock.Anything).Return(nil)
//
// 	result, err := uc.Create(context.Background(), domain.FullURL("https://example.com"))
// 	require.NoError(t, err)
// 	require.NotNil(t, result)
// 	repo.AssertCalled(t, "Save", mock.Anything, link.ShortCode())
// }
