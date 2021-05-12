package internal

import (
	"context"
	"fmt"
	"os"

	redigo "github.com/go-redis/redis/v8"
)

var RedisClient *redigo.Client

func Setup() {
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "127.0.0.1"
	}
	redisPort := os.Getenv("REDIS_PORT")
	if redisPort == "" {
		redisPort = "6379"
	}
	RedisClient = redigo.NewClient(&redigo.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		PoolSize: 6,
	})

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}
