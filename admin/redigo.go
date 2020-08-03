package img2webp

import (
	redigo "github.com/go-redis/redis/v8"
)

var RedisClient *redigo.Client

func Setup() {
	RedisClient = redigo.NewClient(&redigo.Options{
		Addr:     "127.0.0.1:6379",
		PoolSize: 6,
	})
}
