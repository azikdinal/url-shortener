package router

import (
	linkHandler "shorten/internal/transport/link/rest"

	"github.com/gin-gonic/gin"
)

func NewEngine(h *linkHandler.LinkHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/links", h.Create)
	r.GET("/links", h.Get)

	return r
}
