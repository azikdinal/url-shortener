package transport

import (
	"shorten/internal/config"
	"shorten/internal/router"
	linkHandler "shorten/internal/transport/link/rest"
)

func BuildHTTPServer(uc linkHandler.LinkUseCase, cfg *config.Config) *router.Router {
	h := linkHandler.New(uc, cfg.PublicHost)
	return router.New(h)
}
