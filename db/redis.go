package db

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
	AppSetting "github.com/thangsuperman/bee-happy/config"
)

var RedisClient = getRedisClient()

func getRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     AppSetting.Envs.RedisAddress,
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Println("Failed to connect to Redis: ", err)
		return nil
	}
	log.Println("Successfully connected to Redis")
	return client
}
