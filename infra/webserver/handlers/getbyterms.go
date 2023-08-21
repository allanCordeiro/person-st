package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/AllanCordeiro/person-st/application/usecases/person"
)

func (h *PersonHandler) GetByTerms(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
	defer cancel()
	input := person.SearchByTermRequestInput{Term: r.URL.Query().Get("t")}
	if input.Term == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	usecase := person.NewSearchByTermUseCase(h.PersonGateway)
	personOutput, err := usecase.Execute(ctx, input)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&personOutput)
}
