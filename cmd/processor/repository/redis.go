package repository

import (
	"github.com/go-redis/redis"
)

func getRedisConnection() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       10,
	})

	return client
}
