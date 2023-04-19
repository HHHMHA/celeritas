package celeritas

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (c *Celeritas) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Recoverer,
		c.SessionLoad,
	)
	if c.Debug {
		mux.Use(middleware.Logger)
	}

	return mux
}
