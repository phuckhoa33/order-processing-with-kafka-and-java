package config

import (
	"os"
)

type RedisConfig struct {
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDb       string
}

func LoadRedisConfig() RedisConfig {
	return RedisConfig{
		RedisHost:     os.Getenv("REDIS_HOST"),
		RedisPort:     os.Getenv("REDIS_PORT"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisDb:       os.Getenv("REDIS_DB"),
	}
}
