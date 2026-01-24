package router

import (
	"context"
	"net/http"

	"shorten/internal/handler"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
	server *http.Server
}

func New(h *handler.Handler) *Router {
	r := gin.Default()

	r.POST("/links", h.Create)
	r.GET("/links/:code", h.GetByCode)

	return &Router{
		engine: r,
	}
}

func (r *Router) Run(addr string) error {
	r.server = &http.Server{
		Addr:    addr,
		Handler: r.engine,
	}
	return r.server.ListenAndServe()
}

func (r *Router) Shutdown(ctx context.Context) error {
	return r.server.Shutdown(ctx)
}
