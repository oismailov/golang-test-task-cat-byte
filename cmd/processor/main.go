package main

import "twitch_chat_analysis/cmd/processor/repository"

func main() {
	repository.ConsumeMessages()
}
