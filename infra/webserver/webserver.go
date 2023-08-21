package webserver

import (
	"log"
	"net/http"
	"os"

	"github.com/AllanCordeiro/person-st/application/gateway"
	"github.com/AllanCordeiro/person-st/infra/cache"
	"github.com/AllanCordeiro/person-st/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
)

func Serve(personGateway gateway.PersonGateway, cache cache.Cache) {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8181"
	}

	handlers := handlers.NewPersonHandler(personGateway, cache)
	r := chi.NewRouter()
	r.Route("/pessoas", func(r chi.Router) {
		r.Post("/", handlers.CreatePerson)
		r.Get("/", handlers.GetByTerms)
		r.Get("/{personID}", handlers.GetPersonById)
	})
	go r.Route("/contagem-pessoas", func(r chi.Router) {
		r.Get("/", handlers.GetTotal)
	})

	log.Fatal(http.ListenAndServe(":"+port, r))
}
