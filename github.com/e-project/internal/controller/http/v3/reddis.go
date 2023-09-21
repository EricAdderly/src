package v3

import (
	"github.com/go-redis/redis"
)

func dial() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
