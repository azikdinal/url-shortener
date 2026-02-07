package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type getResponse struct {
	FullURL string `json:"full_url"`
}

func (h *LinkHandler) GetByCode(c *gin.Context) {
	ctx := c.Request.Context()

	shortCode := c.Param("short_code")

	fullURL, err := h.linkUseCase.GetFullURL(ctx, shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "link not found"})
		return
	}

	res := getResponse{
		FullURL: fullURL,
	}
	c.JSON(http.StatusOK, res)
}
