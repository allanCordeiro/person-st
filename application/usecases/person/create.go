package person

import (
	"github.com/AllanCordeiro/person-st/application/domain"
	"github.com/AllanCordeiro/person-st/application/gateway"
)

type CreatePersonUseCase struct {
	PersonGateway gateway.PersonGateway
}

func NewCreatePersonUseCase(personGateway gateway.PersonGateway) *CreatePersonUseCase {
	return &CreatePersonUseCase{
		PersonGateway: personGateway,
	}
}

type CreatePersonRequestInput struct {
	NickName  string   `json:"apelido"`
	Name      string   `json:"nome"`
	BirthDate string   `json:"nascimento"`
	StackList []string `json:"stack"`
}

type CreatePersonRequestOutput struct {
	ID string
}

func (u *CreatePersonUseCase) Execute(input CreatePersonRequestInput) (*CreatePersonRequestOutput, error) {
	newPerson, err := domain.CreatePerson(input.NickName, input.Name, input.BirthDate)
	if err != nil {
		return nil, err
	}
	stackList, err := generateStackList(input.StackList)
	if err != nil {
		return nil, err
	}
	if len(stackList.GetStacks()) > 0 {
		newPerson.AddStackList(stackList.GetStacks())
	}
	err = u.PersonGateway.Save(*newPerson)
	if err != nil {
		return nil, err
	}
	return &CreatePersonRequestOutput{ID: newPerson.Id}, nil
}

func generateStackList(stackList []string) (*domain.StackList, error) {
	list := domain.StackList{}
	for _, stackName := range stackList {
		stack, err := domain.NewStack(stackName)
		if err != nil {
			return nil, err
		}
		list.AddStack(*stack)
	}
	return &list, nil
}
