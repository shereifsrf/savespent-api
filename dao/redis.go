package dao

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

func connectRedis() {
	fmt.Println("Connecting to Redis...")

	// connect to redis, throw error if connection fails
	Redis = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	_, err := Redis.Ping(Redis.Context()).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Redis")
}
