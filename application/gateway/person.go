package gateway

import "github.com/AllanCordeiro/person-st/application/domain"

type PersonGateway interface {
	Save(domain.Person) error
	GetByID(id string) (*domain.Person, error)
	GetByTerms(term string) ([]domain.Person, error)
	GetTotal() (*int64, error)
}

//https://arctype.com/blog/postgres-full-text-search/
