package application

import (
	"twitch_chat_analysis/cmd/api/model"
	"twitch_chat_analysis/cmd/api/repository"
)

func SaveMessage(message *model.Message) error {
	return repository.SaveMessage(message)
}
