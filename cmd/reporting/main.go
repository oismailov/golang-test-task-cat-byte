package main

import (
	"log"

	"twitch_chat_analysis/cmd/reporting/config"
	r "twitch_chat_analysis/cmd/reporting/router"
)

func main() {
	config.LoadConfig()

	router := r.GetRouter()

	err := router.Run(r.GetPort())

	if err != nil {
		log.Fatal("unable to start service")
	}
}
