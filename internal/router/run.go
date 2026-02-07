package router

import (
	"net/http"
)

func (r *Router) Run(addr string) error {
	r.server = &http.Server{
		Addr:    addr,
		Handler: r.engine,
	}
	return r.server.ListenAndServe()
}
