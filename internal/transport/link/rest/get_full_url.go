package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type getResponse struct {
	FullURL string `json:"full_url"`
}
type getRequest struct {
	ShortCode string `json:"short_code"`
}

// Обработчик GET запроса
func (h *LinkHandler) GetFullURL(c *gin.Context) {
	ctx := c.Request.Context()

	var req getRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неправильное значение full_url"})
		return
	}

	fullURL, err := h.linkUseCase.GetFullURL(ctx, req.ShortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "link not found"})
		return
	}

	res := getResponse{
		FullURL: fullURL,
	}
	c.JSON(http.StatusOK, res)
}
