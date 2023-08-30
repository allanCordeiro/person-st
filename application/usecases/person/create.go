package person

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/AllanCordeiro/person-st/application/domain"
	"github.com/AllanCordeiro/person-st/infra/cache"
	"github.com/AllanCordeiro/person-st/infra/queue"
)

type CreatePersonUseCase struct {
	Cache cache.Cache
	Queue queue.RabbitMQImpl
}

func NewCreatePersonUseCase(cache cache.Cache, queue queue.RabbitMQImpl) *CreatePersonUseCase {
	return &CreatePersonUseCase{
		Cache: cache,
		Queue: queue,
	}
}

type CreatePersonRequestInput struct {
	NickName  string   `json:"apelido"`
	Name      string   `json:"nome"`
	BirthDate string   `json:"nascimento"`
	StackList []string `json:"stack"`
}

func (c *CreatePersonRequestInput) String() (string, error) {
	data, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(data), nil
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

		message, err := input.String()
		if err != nil {
			panic(err)
		}

		go u.Queue.Publish("person.created", message)

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
