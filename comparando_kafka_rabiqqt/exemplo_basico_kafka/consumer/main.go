package main

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	bootstrapServers := os.Getenv("KAFKA_BOOTSTRAP_SERVERS")

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}

	defer c.Close()

	c.SubscribeTopics([]string{"myTopic"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
