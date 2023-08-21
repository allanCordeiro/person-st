package person

import (
	"errors"
	"log"

	"github.com/AllanCordeiro/person-st/application/domain"
	"github.com/AllanCordeiro/person-st/application/gateway"
	"github.com/AllanCordeiro/person-st/infra/cache"
)

type CreatePersonUseCase struct {
	PersonGateway gateway.PersonGateway
	Cache         cache.Cache
}

func NewCreatePersonUseCase(personGateway gateway.PersonGateway, cache cache.Cache) *CreatePersonUseCase {
	return &CreatePersonUseCase{
		PersonGateway: personGateway,
		Cache:         cache,
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
	err := u.Cache.GetNickname(input.NickName)
	if err != nil && err.Error() == "redigo: nil returned" {
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

		u.setCache(*newPerson)
		err = u.Cache.SetNickname(newPerson.NickName)
		if err != nil {
			log.Println("erro ao add no cache " + err.Error())
		}
		return &CreatePersonRequestOutput{ID: newPerson.Id}, nil
	}
	return nil, errors.New("nickname already exists")
}

func (u *CreatePersonUseCase) setCache(person domain.Person) {
	err := u.Cache.Set(person)
	if err != nil {
		log.Println("error to set cache: " + err.Error())
	}
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
