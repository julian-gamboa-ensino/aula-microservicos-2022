package main

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	bootstrapServers := os.Getenv("KAFKA_BOOTSTRAP_SERVERS")

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers})
	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
	}

	defer p.Close()

	topic := "myTopic"
	for _, word := range []string{"Hello", "World"} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}

	p.Flush(15 * 1000)
}
