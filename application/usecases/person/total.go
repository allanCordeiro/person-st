package person

import (
	"log"

	"github.com/AllanCordeiro/person-st/application/gateway"
)

type GetTotalPersonUseCase struct {
	PersonGateway gateway.PersonGateway
}

func NewGetTotalPersonUseCase(personGateway gateway.PersonGateway) *GetTotalPersonUseCase {
	return &GetTotalPersonUseCase{
		PersonGateway: personGateway,
	}
}

type GetTotalRequestOutput struct {
	Value int64 `json:"total"`
}

func (u *GetTotalPersonUseCase) Execute() *GetTotalRequestOutput {
	total, err := u.PersonGateway.GetTotal()
	if err != nil {
		log.Println(err)
		return nil
	}

	return &GetTotalRequestOutput{Value: *total}
}
