package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type createRequest struct {
	FullURL string `json:"full_url" binding:"required,url"`
}
type createResponse struct {
	ShortURL string `json:"short_url"`
}

// Обработчик POST запроса
func (h *LinkHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var req createRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON format"})
		return
	}

	shortURL, err := h.linkUseCase.Create(ctx, req.FullURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot create link"})
		return
	}

	res := createResponse{
		ShortURL: shortURL,
	}
	c.JSON(http.StatusOK, res)
}
