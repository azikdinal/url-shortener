package router

import (
	"context"
)

func (r *Router) Shutdown(ctx context.Context) error {
	return r.server.Shutdown(ctx)
}
