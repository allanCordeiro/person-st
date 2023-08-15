package webserver

import (
	"log"
	"net/http"

	"github.com/AllanCordeiro/person-st/application/gateway"
	"github.com/AllanCordeiro/person-st/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
)

func Serve(personGateway gateway.PersonGateway) {

	handlers := handlers.NewPersonHandler(personGateway)
	r := chi.NewRouter()
	r.Route("/pessoas", func(r chi.Router) {
		r.Post("/", handlers.CreatePerson)
		r.Get("/", handlers.GetByTerms)
		r.Get("/{personID}", handlers.GetPersonById)
	})
	r.Route("/contagem-pessoas", func(r chi.Router) {
		r.Get("/", handlers.GetTotal)
	})

	log.Fatal(http.ListenAndServe(":8181", r))
}
