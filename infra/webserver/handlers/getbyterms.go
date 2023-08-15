package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AllanCordeiro/person-st/application/usecases/person"
)

func (h *PersonHandler) GetByTerms(w http.ResponseWriter, r *http.Request) {
	input := person.SearchByTermRequestInput{Term: r.URL.Query().Get("t")}
	if input.Term == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	usecase := person.NewSearchByTermUseCase(h.PersonGateway)
	personOutput, err := usecase.Execute(input)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&personOutput)
}
