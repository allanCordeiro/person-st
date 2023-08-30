package domain

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Person struct {
	Id        string
	NickName  string
	Name      string
	BirthDate time.Time
	StackList []Stack
}

func CreatePerson(nickname string, name string, birthdate string) (*Person, error) {
	parsedBirth, err := time.Parse("2006-01-02", birthdate)
	if err != nil {
		return nil, errors.New("invalid birth date")
	}

	person := &Person{
		Id:        uuid.New().String(),
		NickName:  nickname,
		Name:      name,
		BirthDate: parsedBirth,
	}
	err = person.Validate()
	if err != nil {
		return nil, err
	}
	return person, nil
}

func (p *Person) AddStackList(stackList []Stack) {
	p.StackList = append(p.StackList, stackList...)
}

func (p *Person) GetStackListToString() []string {
	list := []string{}
	for _, stack := range p.StackList {
		list = append(list, stack.Name)
	}
	return list
}

func (p *Person) Validate() error {

	if p.NickName == "" {
		return errors.New("nickname is null")
	}
	if len(p.NickName) > 32 {
		return errors.New("nickname is greater than 32 chars")
	}
	if p.Name == "" {
		return errors.New("name is null")
	}
	if len(p.Name) > 100 {
		return errors.New("name is greater than 100 chars")
	}
	return nil
}

func BuildPerson(id string, nickname string, name string, birthdate time.Time, stack []Stack) (*Person, error) {
	return &Person{
		Id:        id,
		NickName:  nickname,
		Name:      name,
		BirthDate: birthdate,
		StackList: stack,
	}, nil
}

func (p *Person) String() (string, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
