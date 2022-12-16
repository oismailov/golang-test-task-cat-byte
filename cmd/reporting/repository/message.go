package repository

import (
	"fmt"

	"twitch_chat_analysis/cmd/reporting/model"

	"github.com/wagslane/go-rabbitmq"
)

func SaveMessage(message *model.Message) error {
	connection := GetPublisherConnection()
	err := connection.Publish(
		[]byte(fmt.Sprintf("%v", message)),
		[]string{"my_routing_key"},
		rabbitmq.WithPublishOptionsContentType("application/json"),
	)
	if err != nil {
		return err
	}

	return nil
}
