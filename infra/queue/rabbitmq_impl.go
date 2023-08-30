package queue

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQImpl struct {
	Channel *amqp.Channel
}

func NewRabbitMQImpl(channel *amqp.Channel) *RabbitMQImpl {
	return &RabbitMQImpl{Channel: channel}
}

func (r *RabbitMQImpl) Publish(exchange, body string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.Channel.PublishWithContext(
		ctx,
		exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		return err
	}
	return nil
}
