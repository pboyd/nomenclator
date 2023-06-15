package httpserver

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/pboyd/nomenclator/api/internal/domain"
)

// NewHandler returns a new Handler.
func NewHandler(services *domain.Bundle) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {
		r.Route("/people", newPeopleController(services))
	})

	return r
}
