package router

import (
	"fmt"

	"twitch_chat_analysis/cmd/reporting/config"
	"twitch_chat_analysis/cmd/reporting/handler"
	"twitch_chat_analysis/cmd/reporting/middleware"
	"twitch_chat_analysis/cmd/reporting/pkg/util"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.JSONMiddleware(), middleware.CORSMiddleware())
	r.POST("/api/message/list", handler.ListMessages)

	r.NoRoute(func(c *gin.Context) {
		util.ResponseNotFound(c, "Page not found")
	})

	return r
}

func GetPort() string {
	return fmt.Sprintf(":%d", config.Cfg.ServerConfigurations.Port)
}
