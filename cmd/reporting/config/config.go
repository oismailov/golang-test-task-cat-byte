package config

import (
	"encoding/json"
	"log"
	"os"
)

var Cfg = &Config{}

type Config struct {
	RabbitMQ             RabbitMQSettings
	ServerConfigurations ServerConfigurations
}

type RabbitMQSettings struct {
	User     string
	Password string
}

type ServerConfigurations struct {
	Port int
}

func LoadConfig() {
	loadBaseConfig("config/config.json")
}

func loadBaseConfig(path string) {
	configFile, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(configFile)
	configuration := Config{}
	err = decoder.Decode(&configuration)

	if err != nil {
		log.Fatal("error:", err)
	}

	Cfg = &configuration
}
