package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBHost                 string
	DBName                 string
	RedisAddress           string
	JWTSecret              string
	JWTExpirationInSeconds int64
	RabbitMQUsername       string
	RabbitMQPassword       string
	RabbitMQHost           string
	RabbitMQVirtualHost    string
	RabbitMQPort           string
	AWSRegion              string
	AWSBucketAccessKey     string
	AWSBucketSecretKey     string
	AWSBucketName          string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost:             getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                   getEnv("PORT", "3001"),
		DBUser:                 getEnv("DB_USER", "root"),
		DBPassword:             getEnv("DB_PASSWORD", "hellodb"),
		DBAddress:              fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:                 getEnv("DB_NAME", "dev"),
		RedisAddress:           fmt.Sprintf("%s:%s", getEnv("REDIS_HOST", "127.0.0.1"), getEnv("REDIS_PORT", "6379")),
		RabbitMQUsername:       getEnv("RABBITMQ_USERNAME", "guest"),
		RabbitMQPassword:       getEnv("RABBITMQ_PASSWORD", "guest"),
		RabbitMQHost:           getEnv("RABBITMQ_HOST", "localhost"),
		RabbitMQVirtualHost:    getEnv("RABBITMQ_VIRTUAL_HOST", "customers"),
		RabbitMQPort:           getEnv("RABBITMQ_PORT", "secret123"),
		JWTSecret:              getEnv("JWT_SECRET", "secret123"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 3600*24*7),
		AWSRegion:              getEnv("AWS_REGION", "ap-southeast-1"),
		AWSBucketAccessKey:     getEnv("AWS_BUCKET_ACCESS_KEY", "ap-southeast-1"),
		AWSBucketSecretKey:     getEnv("AWS_BUCKET_SECRET_KEY", "ap-southeast-1"),
		AWSBucketName:          getEnv("AWS_BUCKET_NAME", "bee_happy_bucket"),
	}
}

func getEnv(key string, fallback string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
