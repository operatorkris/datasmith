package data

import (
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func StartKafkaConsumer(brokers []string, topic string) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: "datasmith-group",
	})

	for {
		msg, err := r.ReadMessage(nil)
		if err != nil {
			log.Fatalf("Error reading Kafka message: %v", err)
		}
		fmt.Printf("Received message: %s\n", string(msg.Value))
	}
}
