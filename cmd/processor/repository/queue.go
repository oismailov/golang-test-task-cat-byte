package repository

import (
	"encoding/json"
	"fmt"
	"log"

	"twitch_chat_analysis/cmd/api/model"

	"github.com/wagslane/go-rabbitmq"
)

func consumeMessage() {
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

	consumer, err := rabbitmq.NewConsumer(
		conn,
		processMessage,
		"my_queue",
		rabbitmq.WithConsumerOptionsRoutingKey("my_routing_key"),
	)

	if err != nil {
		fmt.Printf("rabbitmq connection error: %v", err)
		return
	}
	defer consumer.Close()
}

func processMessage(d rabbitmq.Delivery) rabbitmq.Action {
	redis := getRedisConnection()
	var messages []model.Message
	var message model.Message

	val, err := redis.Get("messages").Result()
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal([]byte(val), &messages)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(d.Body, &message)
	if err != nil {
		log.Fatal(err)
	}

	messages = append(messages, message)

	payload, _ := json.Marshal(messages)
	redis.Set("messages", payload, 0)

	return rabbitmq.Ack
}
