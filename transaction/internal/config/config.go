package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	App   AppConfig
	DB    DBConfig
	Auth  AuthConfig
	Mail  MailConfig
	Redis RedisConfig
	Kafka KafkaConfig
}

func NewConfig() *Config {

	// Load environment variables and check error
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Get env failed", err)
	}

	env := &Config{
		App:   LoadAppConfig(),
		DB:    LoadDBConfig(),
		Auth:  LoadAuthConfig(),
		Mail:  LoadMailConfig(),
		Redis: LoadRedisConfig(),
		Kafka: LoadKafkaConfig(),
	}

	if env.App.AppEnv == "development" {
		log.Println("The App is running in development app")
	}

	return env
}
