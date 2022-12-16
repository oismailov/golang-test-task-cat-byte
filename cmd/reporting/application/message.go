package application

import (
	"twitch_chat_analysis/cmd/reporting/model"
	"twitch_chat_analysis/cmd/reporting/repository"
)

func SaveMessage(message *model.Message) error {
	return repository.SaveMessage(message)
}
