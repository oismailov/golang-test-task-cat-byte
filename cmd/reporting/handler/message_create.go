package handler

import (
	"encoding/json"
	"log"

	"twitch_chat_analysis/cmd/reporting/model"
	"twitch_chat_analysis/cmd/reporting/pkg/util"
	"twitch_chat_analysis/cmd/reporting/repository"

	"github.com/gin-gonic/gin"
)

func ListMessages(c *gin.Context) {
	redis := repository.GetRedisConnection()
	var messages []model.Message
	val, err := redis.Get("messages").Result()
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal([]byte(val), &messages)
	if err != nil {
		log.Fatal(err)
	}

	util.ResponseCreated(c, messages)
}
