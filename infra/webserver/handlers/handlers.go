package handlers

import (
	"github.com/AllanCordeiro/person-st/application/gateway"
	"github.com/AllanCordeiro/person-st/infra/cache"
)

type PersonHandler struct {
	PersonGateway gateway.PersonGateway
	Cache         cache.Cache
}

func NewPersonHandler(db gateway.PersonGateway, cache cache.Cache) *PersonHandler {
	return &PersonHandler{
		PersonGateway: db,
		Cache:         cache,
	}
}
