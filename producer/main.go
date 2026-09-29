package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

type Event struct {
	EventID   string `json:"event_id"`
	EventType string `json:"event_type"`
	Sequence  int    `json:"sequence"`
	Timestamp int64  `json:"timestamp"`
}

func main() {
	broker := getenv("KAFKA_BOOTSTRAP_SERVERS", "localhost:9092")
	topic := getenv("KAFKA_TOPIC", "events")

	writer := &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    topic,
		Balancer: &kafka.Hash{},
	}
	defer writer.Close()

	for i := 1; i <= 10; i++ {
		event := Event{
			EventID:   fmt.Sprintf("event-%d", i),
			EventType: "platform.event",
			Sequence:  i,
			Timestamp: time.Now().Unix(),
		}

		value, err := json.Marshal(event)
		if err != nil {
			log.Fatal(err)
		}

		err = writer.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(event.EventID),
			Value: value,
		})
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("published event=%s", event.EventID)
	}
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
