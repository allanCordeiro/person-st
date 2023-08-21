package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AllanCordeiro/person-st/application/usecases/person"
	"github.com/go-chi/chi/v5"
)

func (h *PersonHandler) GetPersonById(w http.ResponseWriter, r *http.Request) {
	input := person.GetByIdRequestInput{ID: chi.URLParam(r, "personID")}
	usecase := person.NewGetPersonByIdUseCase(h.PersonGateway, h.Cache)
	personOutput, err := usecase.Execute(input)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	output := person.GetByIdRequestOutput{
		ID:        personOutput.ID,
		NickName:  personOutput.NickName,
		Name:      personOutput.Name,
		BirthDate: personOutput.BirthDate,
		StackList: personOutput.StackList,
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&output)
}
