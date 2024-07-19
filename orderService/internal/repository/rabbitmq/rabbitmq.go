package rabbitmq

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"orderService/internal/domain"
)

type RabbitMQClient interface {
	SendMessage(order *domain.Order) error
}

// RabbitMQClientImpl is a concrete implementation of RabbitMQClient.
type RabbitMQClientImpl struct {
	channel *amqp.Channel
	queue   amqp.Queue
}

// NewRabbitMQClient creates a new instance of RabbitMQClientImpl.
func NewRabbitMQClient(channel *amqp.Channel, queueName string) (*RabbitMQClientImpl, error) {
	q, err := channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMQClientImpl{channel: channel, queue: q}, nil
}

// SendMessage sends a message to RabbitMQ.
func (c *RabbitMQClientImpl) SendMessage(order *domain.Order) error {
	body, err := json.Marshal(order)
	if err != nil {
		return err
	}

	return c.channel.Publish(
		"",
		c.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
