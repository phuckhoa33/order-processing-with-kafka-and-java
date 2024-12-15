package config

import "os"

type KafkaConfig struct {
	KafkaBroker string
	KafkaTopic string
}

// Kafka struct config
func LoadKafkaConfig() KafkaConfig {
	return KafkaConfig{
		KafkaBroker: os.Getenv("KAFKA_BROKER"),
		KafkaTopic: os.Getenv("KAFKA_TOPIC"),
	}
}