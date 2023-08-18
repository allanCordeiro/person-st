package cache

import "github.com/AllanCordeiro/person-st/application/domain"

type Cache interface {
	Get(string) (*domain.Person, error)
	Set(domain.Person) error
}
