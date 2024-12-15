package config

import (
	"log"
	"os"
	"strconv"
)

type AppConfig struct {
	SwaggerEnable  string
	AppName        string
	AppDescription string
	AppVersion     string
	AppEnv         string
	AppPort        string
	AppHost        string
	AppApiDebug    string
	AppApiPrefix   string
	PingTimeout    int
}

func LoadAppConfig() AppConfig {
	// Get ping timeout
	pingTimeout, err := strconv.Atoi(os.Getenv("PING_TIMEOUT"))
	if err != nil {
		log.Fatal(err)
	}

	return AppConfig{
		SwaggerEnable:  os.Getenv("SWAGGER_ENABLE"),
		AppName:        os.Getenv("APP_NAME"),
		AppDescription: os.Getenv("APP_DESCRIPTION"),
		AppVersion:     os.Getenv("APP_VERSION"),
		AppEnv:         os.Getenv("APP_ENV"),
		AppPort:        os.Getenv("APP_PORT"),
		AppHost:        os.Getenv("APP_HOST"),
		AppApiDebug:    os.Getenv("APP_API_DEBUG"),
		AppApiPrefix:   os.Getenv("APP_API_PREFIX"),
		PingTimeout:    pingTimeout,
	}
}
