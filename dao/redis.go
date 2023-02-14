package dao

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

var (
	redisClient *redis.Client
)

func init() {
	fmt.Println("Redis INIT()")
	redisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL") + ":6379",
		Password: "",
		DB:       0,
	})
}

func Get(ctx context.Context, key string) string {
	value, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		fmt.Printf("[Get] redisClient.Get(key: %s), err: %v", key, err)
		return ""
	}

	return value
}

func Set(ctx context.Context, key string, value string) error {
	err := redisClient.Set(ctx, key, value, 0).Err()
	if err != nil {
		fmt.Printf("[Set] redisClient.Set(key: %s, value: %s), err: %v", key, value, err)
		return err
	}

	return nil
}
