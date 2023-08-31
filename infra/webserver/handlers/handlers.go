package handlers

import (
	"github.com/AllanCordeiro/person-st/application/gateway"
	"github.com/AllanCordeiro/person-st/infra/cache"
	"github.com/AllanCordeiro/person-st/infra/queue"
)

type PersonHandler struct {
	PersonGateway gateway.PersonGateway
	Cache         cache.Cache
	Queue         *queue.RabbitMQImpl
}

func NewPersonHandler(db gateway.PersonGateway, cache cache.Cache, queue *queue.RabbitMQImpl) *PersonHandler {
	return &PersonHandler{
		PersonGateway: db,
		Cache:         cache,
		Queue:         queue,
	}
}
