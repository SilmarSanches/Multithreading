package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/silmarsanches/multithreading/internal/web/middlewares"
	"net/http"
)

func NewServer(cepRouters http.Handler) *http.Server {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middlewares.MiddlewareLog)

	r.Mount("/buscacep", cepRouters)

	return &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
}
