package router

import (
	"net/http"

	linkHandler "shorten/internal/transport/link/rest"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
	server *http.Server
}

func New(h *linkHandler.LinkHandler) *Router {
	r := gin.Default()

	r.POST("/links", h.Create)
	r.GET("/links/:code", h.GetByCode)

	return &Router{
		engine: r,
	}
}
