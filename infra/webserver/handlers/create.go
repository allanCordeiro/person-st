package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AllanCordeiro/person-st/application/usecases/person"
)

type RequestCreateOutput struct {
	value string
}

func (h *PersonHandler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	var input person.CreatePersonRequestInput

	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	usecase := person.NewCreatePersonUseCase(h.Cache, *h.Queue)
	personOutput, err := usecase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	output := &RequestCreateOutput{
		value: "/pessoas/" + personOutput.ID,
	}

	w.Header().Add("Location", output.value)
	w.WriteHeader(http.StatusCreated)
}
