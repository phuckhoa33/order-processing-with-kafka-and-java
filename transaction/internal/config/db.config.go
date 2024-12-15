package config

import (
	"os"
)

// PostgresDBConfig struct
type PostgresDBConfig struct {
	PostgresDatabaseHost     string
	PostgresDatabasePort     string
	PostgresDatabaseName     string
	PostgresDatabaseUser     string
	PostgresDatabasePassword string
	PostgresDatabaseDebug    string
}

// DBConfig DbConfig struct
type DBConfig struct {
	PostgresConfig PostgresDBConfig
}

// LoadDBConfig function
func LoadDBConfig() DBConfig {
	return DBConfig{
		PostgresConfig: PostgresDBConfig{
			PostgresDatabaseHost:     os.Getenv("POSTGRES_DATABASE_HOST"),
			PostgresDatabasePort:     os.Getenv("POSTGRES_DATABASE_PORT"),
			PostgresDatabaseName:     os.Getenv("POSTGRES_DATABASE_NAME"),
			PostgresDatabaseUser:     os.Getenv("POSTGRES_DATABASE_USER"),
			PostgresDatabasePassword: os.Getenv("POSTGRES_DATABASE_PASSWORD"),
			PostgresDatabaseDebug:    os.Getenv("POSTGRES_DATABASE_DEBUG"),
		},
	}
}
