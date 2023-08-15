package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AllanCordeiro/person-st/application/usecases/person"
)

func (h *PersonHandler) GetTotal(w http.ResponseWriter, r *http.Request) {
	usecase := person.NewGetTotalPersonUseCase(h.PersonGateway)
	total := usecase.Execute()
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&total)
}
