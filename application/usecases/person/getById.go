package person

import (
	"github.com/AllanCordeiro/person-st/application/gateway"
)

type GetPersonByIdUseCase struct {
	PersonGateway gateway.PersonGateway
}

func NewGetPersonByIdUseCase(personGateway gateway.PersonGateway) *GetPersonByIdUseCase {
	return &GetPersonByIdUseCase{
		PersonGateway: personGateway,
	}
}

type GetByIdRequestInput struct {
	ID string
}

type GetByIdRequestOutput struct {
	ID        string   `json:"id"`
	NickName  string   `json:"apelido"`
	Name      string   `json:"nome"`
	BirthDate string   `json:"nascimento"`
	StackList []string `json:"stack"`
}

func (u *GetPersonByIdUseCase) Execute(input GetByIdRequestInput) (*GetByIdRequestOutput, error) {
	person, err := u.PersonGateway.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	stackList := person.GetStackListToString()

	return &GetByIdRequestOutput{
		ID:        person.Id,
		NickName:  person.NickName,
		Name:      person.Name,
		BirthDate: person.BirthDate.Format("2006-01-02"),
		StackList: ShouldSendStackList(stackList),
	}, nil
}

func ShouldSendStackList(list []string) []string {
	if len(list) == 0 {
		return nil
	}
	return list
}
