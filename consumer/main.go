package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

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
	groupID := getenv("KAFKA_GROUP_ID", "processing-workers")

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{broker},
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 1,
		MaxBytes: 10e6,
	})
	defer reader.Close()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	log.Printf("consuming topic=%s group=%s", topic, groupID)

	for {
		message, err := reader.ReadMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				return
			}
			log.Printf("read error: %v", err)
			continue
		}

		var event Event
		if err := json.Unmarshal(message.Value, &event); err != nil {
			log.Printf("invalid event offset=%d: %v", message.Offset, err)
			continue
		}

		log.Printf(
			"received event=%s partition=%d offset=%d",
			event.EventID,
			message.Partition,
			message.Offset,
		)
	}
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
