package router

import (
	"fmt"

	"twitch_chat_analysis/cmd/api/config"
	"twitch_chat_analysis/cmd/api/handler"
	"twitch_chat_analysis/cmd/api/middleware"
	"twitch_chat_analysis/cmd/api/pkg/util"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.JSONMiddleware(), middleware.CORSMiddleware())
	r.POST("/api/message", handler.CreateMessage)

	r.NoRoute(func(c *gin.Context) {
		util.ResponseNotFound(c, "Page not found")
	})

	return r
}

func GetPort() string {
	return fmt.Sprintf(":%d", config.Cfg.ServerConfigurations.Port)
}
