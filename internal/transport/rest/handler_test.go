package rest_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"shorten/internal/domain"
	"shorten/internal/transport/rest"
	"shorten/internal/usecase/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRESTHandler_Create(t *testing.T) {
	host := "http://sh.by/"

	mockUC := new(mocks.Usecase)
	handler := rest.NewHandler(mockUC, host)

	r := gin.New()
	r.POST("/links", handler.Create)

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

func TestRESTHandler_GetByCode(t *testing.T) {
	host := "http://sh.by"

	mockUC := new(mocks.Usecase)
	handler := rest.NewHandler(mockUC, host)

	r := gin.New()
	r.GET("/links/:code", handler.GetByCode)

	testSC := "dmihf97_v_"
	req := httptest.NewRequest("GET", "/links/"+testSC, nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	testFU, _ := domain.NewFullURL("https://example.com")
	mockUC.On("GetByCode", mock.Anything, domain.ShortCode(testSC))

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	require.Contains(t, w.Body.String(), testFU)
}
