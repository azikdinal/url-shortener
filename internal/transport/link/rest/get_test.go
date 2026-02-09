package rest

import (
	"encoding/json"
	"net/http"
	"shorten/internal/transport/link/mocks"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestLinkHandler_Get_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	useCase := new(mocks.LinkUseCase)
	useCase.
		On("GetFullURL", mock.Anything, "abc123").
		Return("https://example.com", nil).
		Once()

	handler := &LinkHandler{
		linkUseCase: useCase,
	}

	router := gin.New()
	router.POST("/get", handler.Get)

	body := `{"short_code":"abc123"}`
	w := performRequest(router, "POST", "/get", strings.NewReader(body))

	require.Equal(t, http.StatusOK, w.Code)

	var resp getResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Equal(t, "https://example.com", resp.FullURL)

	useCase.AssertExpectations(t)
}
