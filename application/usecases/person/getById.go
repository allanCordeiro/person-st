package person

import (
	"log"

	"github.com/AllanCordeiro/person-st/application/domain"
	"github.com/AllanCordeiro/person-st/application/gateway"
	"github.com/AllanCordeiro/person-st/infra/cache"
)

type GetPersonByIdUseCase struct {
	PersonGateway gateway.PersonGateway
	Cache         cache.Cache
}

func NewGetPersonByIdUseCase(personGateway gateway.PersonGateway, cache cache.Cache) *GetPersonByIdUseCase {
	return &GetPersonByIdUseCase{
		PersonGateway: personGateway,
		Cache:         cache,
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
	person := u.getCache(input.ID)
	if person == nil {
		//return nil, errors.New("sql: no rows in result set")
		var err error
		person, err = u.PersonGateway.GetByID(input.ID)
		if err != nil {
			return nil, err
		}
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

func (u *GetPersonByIdUseCase) getCache(id string) *domain.Person {
	person, error := u.Cache.Get(id)
	if error != nil {
		log.Println("error to get cache: " + error.Error())
	}
	return person
}

func ShouldSendStackList(list []string) []string {
	if len(list) == 0 {
		return nil
	}
	return list
}
