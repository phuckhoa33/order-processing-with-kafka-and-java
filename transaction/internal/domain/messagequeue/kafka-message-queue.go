package messagequeue

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer(broker string, topic string) *KafkaProducer {
	return &KafkaProducer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(broker),
			Topic:    topic,
			Balancer: &kafka.Hash{},
		},
	}
}

func (kp *KafkaProducer) PublishMessage(ctx context.Context, key, value string) error {
	msg := kafka.Message{
		Key:   []byte(key),
		Value: []byte(value),
	}
	return kp.writer.WriteMessages(ctx, msg)
}

func (kp *KafkaProducer) Close() {
	kp.writer.Close()
}