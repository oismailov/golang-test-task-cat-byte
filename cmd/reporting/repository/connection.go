package repository

import (
	"fmt"

	"github.com/wagslane/go-rabbitmq"
)

func GetPublisherConnection() *rabbitmq.Publisher {
	conn, err := rabbitmq.NewConn(
		"amqp://user:password@localhost",
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		fmt.Printf("rabbitmq connection error: %v", err)
	}

	defer func() {
		err = conn.Close()
	}()

	publisher, err := rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsLogging,
	)

	if err != nil {
		fmt.Printf("rabbitmq connection error: %v", err)
		return nil
	}
	defer publisher.Close()

	return publisher
}
