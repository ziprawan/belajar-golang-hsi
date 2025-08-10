package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

type Event struct {
	StudentID string `json:"student_id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
}

func main() {
	log.Printf("[StudentService] Started\n")
	log.Printf("[StudentService] Creating writer and event\n")
	writer := kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "student.registered",
	}
	defer writer.Close()

	event := &Event{
		StudentID: "1",
		Name:      "Aziz",
		Status:    "student.registered",
	}

	log.Printf("[StudentService] Converting event to bytes\n")
	eventBytes, err := json.Marshal(event)
	if err != nil {
		log.Fatalf("[StudentService] Failed to convert struct into bytes: %s\n", err.Error())
	}

	log.Printf("[StudentService] Writing message\n")
	err = writer.WriteMessages(
		context.TODO(),
		kafka.Message{
			Key:   []byte("msg-1"),
			Value: eventBytes,
		},
	)
	if err != nil {
		log.Fatalf("[StudentService] Failed to write message: %s\n", err.Error())
	}

	log.Printf("[StudentService] Sent event: %s\n", event.Status)

	log.Printf("[StudentService] Creating reader for failed registration\n")
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "student.registration_failed",
		GroupID: "student-service",
	})

	log.Printf("[StudentService] Starting reader\n")
	ctx, cancel := context.WithTimeout(
		context.Background(),
		60*time.Second,
	)
	defer cancel()

	m, err := reader.ReadMessage(ctx)
	if err != nil {
		log.Printf("[StudentService] Payment success!\n")
		os.Exit(0)
	}

	log.Printf("[StudentService] Received event value: %s\n", string(m.Value))
}
