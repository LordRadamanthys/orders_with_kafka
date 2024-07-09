package kafkaconfig

import (
	"time"

	"github.com/segmentio/kafka-go"
)

var (
	KafkaClient *kafka.Reader
)

func InitKafkaConfig() {
	// Defina o endereço do broker Kafka
	kafkaURL := "localhost:9092"
	topic := "broker-kafka"
	// partition := 0

	// Inicialize um leitor Kafka com a configuração
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaURL},
		Topic:   topic,
		// Partition:     partition,
		RetentionTime: 15 * time.Second,
		GroupID:       "my-group",
		MinBytes:      10e3, // 10KB
		MaxBytes:      10e6, // 10MB

	})

	KafkaClient = r
}
