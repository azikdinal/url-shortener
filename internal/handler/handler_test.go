package handler_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"shorten/internal/domain"
	"shorten/internal/handler"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockUsecase struct {
	mock.Mock
}

func (m *MockUsecase) GetByCode(ctx context.Context, sc domain.ShortCode) (domain.FullURL, error) {
	args := m.Called(ctx, sc)
	return args.Get(0).(domain.FullURL), args.Error(1)
}

func (m *MockUsecase) Create(ctx context.Context, fu domain.FullURL) (*domain.Link, error) {
	args := m.Called(ctx, fu)
	return args.Get(0).(*domain.Link), args.Error(1)
}

func TestHandler_Create(t *testing.T) {
	host := "http://sh.by/"

	mockUC := new(MockUsecase)
	h := handler.New(mockUC, host)

	r := gin.New()
	r.POST("/links", h.Create)

	reqBody := `{"full_url":"https://example.com"}`
	req := httptest.NewRequest("POST", "/links", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	link := domain.NewLink(domain.ShortCode("dfineidmno"), domain.FullURL("https://example.com"))
	mockUC.On("Create", mock.Anything, domain.FullURL("https://example.com")).Return(link, nil)

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	require.Contains(t, w.Body.String(), "https://example.com")
}
