package gateway

import (
	"context"

	"github.com/AllanCordeiro/person-st/application/domain"
)

type PersonGateway interface {
	BulkInsert(people []domain.Person) error
	Save(domain.Person) error
	GetByID(id string) (*domain.Person, error)
	GetByTerms(ctx context.Context, term string) ([]domain.Person, error)
	GetTotal() (*int64, error)
}
