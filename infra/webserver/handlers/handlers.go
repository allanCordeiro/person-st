package handlers

import "github.com/AllanCordeiro/person-st/application/gateway"

type PersonHandler struct {
	PersonGateway gateway.PersonGateway
}

func NewPersonHandler(db gateway.PersonGateway) *PersonHandler {
	return &PersonHandler{PersonGateway: db}
}
