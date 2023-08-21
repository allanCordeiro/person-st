package person

import (
	"context"

	"github.com/AllanCordeiro/person-st/application/gateway"
)

type SearchByTermUseCase struct {
	PersonGateway gateway.PersonGateway
}

func NewSearchByTermUseCase(personGateway gateway.PersonGateway) *SearchByTermUseCase {
	return &SearchByTermUseCase{
		PersonGateway: personGateway,
	}
}

type SearchByTermRequestInput struct {
	Term string
}

func (u *SearchByTermUseCase) Execute(ctx context.Context, input SearchByTermRequestInput) (*[]GetByIdRequestOutput, error) {
	list, err := u.PersonGateway.GetByTerms(ctx, input.Term)
	if err != nil {
		return nil, err
	}

	personList := []GetByIdRequestOutput{}
	for _, person := range list {
		var p GetByIdRequestOutput
		p.ID = person.Id
		p.NickName = person.NickName
		p.Name = person.Name
		p.BirthDate = person.BirthDate.Format("2006-01-02")
		p.StackList = ShouldSendStackList(person.GetStackListToString())
		personList = append(personList, p)
	}
	return &personList, nil
}
