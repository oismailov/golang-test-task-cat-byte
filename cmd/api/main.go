package main

import (
	"log"

	"twitch_chat_analysis/cmd/api/config"
	r "twitch_chat_analysis/cmd/api/router"
)

func main() {
	config.LoadConfig()

	router := r.GetRouter()

	err := router.Run(r.GetPort())

	if err != nil {
		log.Fatal("unable to start service")
	}
}
