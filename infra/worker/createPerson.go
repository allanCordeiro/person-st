package worker

import (
	"encoding/json"
	"log"
	"time"

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
	var peopleBulk []domain.Person
	timer := time.NewTicker(4 * time.Second)

	go p.Queue.Consume("person.created", msgs)

	for {
		select {
		case msg := <-msgs:
			peopleBulk = append(peopleBulk, p.hydrate(msg.Body))
			if len(peopleBulk) >= 500 {
				err := p.PersonGateway.BulkInsert(peopleBulk)
				if err != nil {
					log.Println(err)
					msg.Nack(false, true)
				}
				msg.Ack(true)
				peopleBulk = nil
			}
		case <-timer.C:
			if len(peopleBulk) > 0 {
				err := p.PersonGateway.BulkInsert(peopleBulk)
				if err != nil {
					log.Println(err)
				}
				peopleBulk = nil
			}
		}
	}
}

func (p *CreatePersonWorker) hydrate(person []byte) domain.Person {
	var input domain.Person
	err := json.Unmarshal(person, &input)
	if err != nil {
		log.Printf("error to hydrate message: %v", person)
		return domain.Person{}
	}
	return input
}
