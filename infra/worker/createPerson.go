package worker

import (
	"encoding/json"
	"log"

	"github.com/AllanCordeiro/person-st/application/domain"
	"github.com/AllanCordeiro/person-st/application/gateway"
	"github.com/AllanCordeiro/person-st/infra/queue"
	amqp "github.com/rabbitmq/amqp091-go"
)

type CreatePersonInput struct {
	NickName  string   `json:"apelido"`
	Name      string   `json:"nome"`
	BirthDate string   `json:"nascimento"`
	StackList []string `json:"stack"`
}

type CreatePersonWorker struct {
	PersonGateway gateway.PersonGateway
	Queue         queue.RabbitMQImpl
}

func NewCreatePersonWorker(gateway gateway.PersonGateway, queue queue.RabbitMQImpl) *CreatePersonWorker {
	return &CreatePersonWorker{
		PersonGateway: gateway,
		Queue:         queue,
	}
}

func (p *CreatePersonWorker) Run() {
	msgs := make(chan amqp.Delivery)

	go p.Queue.Consume("person.created", msgs)
	var peopleBulk []domain.Person
	for msg := range msgs {
		peopleBulk = append(peopleBulk, p.hydrate(msg.Body))
		if len(peopleBulk) >= 500 {
			err := p.PersonGateway.BulkInsert(peopleBulk)
			if err != nil {
				log.Println(err)
				msg.Nack(false, true)
			}
			msg.Ack(false)
			peopleBulk = nil
		}
	}
}

func (p *CreatePersonWorker) hydrate(person []byte) domain.Person {
	var input domain.Person
	err := json.Unmarshal(person, &input)
	if err != nil {
		//TODO: temporary. take this off ASAP
		panic(err)
	}
	return input
}
