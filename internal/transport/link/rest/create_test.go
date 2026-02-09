package rest

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"shorten/internal/transport/link/mocks"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func setupRouter(handler *LinkHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/create", handler.Create)
	return r
}

func performRequest(
	r http.Handler,
	method, path string,
	body io.Reader,
) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestLinkHandler_Create_Success(t *testing.T) {
	useCase := new(mocks.LinkUseCase)
	useCase.
		On("Create", mock.Anything, "https://example.com").
		Return("https://short.ly/abc123", nil).
		Once()

	handler := &LinkHandler{
		linkUseCase: useCase,
	}

	router := setupRouter(handler)

	body := `{"full_url":"https://example.com"}`
	w := performRequest(router, "POST", "/create", strings.NewReader(body))

	require.Equal(t, http.StatusOK, w.Code)

	var resp createResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Equal(t, "https://short.ly/abc123", resp.ShortURL)

	useCase.AssertExpectations(t)
}
