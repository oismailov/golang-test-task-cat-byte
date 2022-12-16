package handler

import (
	"twitch_chat_analysis/cmd/api/application"
	"twitch_chat_analysis/cmd/api/model"
	"twitch_chat_analysis/cmd/api/pkg/util"
	validation "twitch_chat_analysis/cmd/api/validiation"

	"github.com/gin-gonic/gin"
)

func CreateMessage(c *gin.Context) {
	message := model.Message{UUID: util.GetUUID()}
	err := c.BindJSON(&message)

	if err != nil {
		util.ResponseBadRequest(c, "invalid request body")
		return
	}

	if err = validation.ValidateCreateMessageRequest(message); err != nil {
		util.ResponseBadRequest(c, err.Error())
		return
	}

	err = application.SaveMessage(&message)
	if err != nil {
		util.ResponseBadRequest(c, "unable to save a message")
		return
	}

	util.ResponseCreated(c, message)
}
