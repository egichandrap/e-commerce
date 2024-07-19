package config

import (
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQ struct {
	Url string
}

func (c RabbitMQ) Connect() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(c.Url)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		return nil, nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		log.Fatalf("Failed to connect to Channel: %v", err)
		return nil, nil, err
	}

	return conn, channel, nil
}
