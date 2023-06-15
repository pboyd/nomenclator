package httpserver

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/pboyd/nomenclator/api/internal/domain"
)

type peopleController struct {
	personService *domain.PersonService
}

func newPeopleController(services *domain.Bundle) func(r chi.Router) {
	c := &peopleController{
		personService: services.PersonService,
	}

	return func(r chi.Router) {
		r.Get("/", c.List)
		r.Post("/", c.Create)
	}
}

func (c *peopleController) List(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (c *peopleController) Create(w http.ResponseWriter, r *http.Request) {
	var person domain.Person
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := c.personService.Create(r.Context(), &person); err != nil {
		log.Printf("failed to create person: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
