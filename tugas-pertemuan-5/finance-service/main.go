package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"

	"github.com/segmentio/kafka-go"
)

type Event struct {
	StudentID string `json:"student_id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
}

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "student.registered",
		GroupID: "academic-service",
	})
	defer reader.Close()

	log.Printf("[FinanceService] Started\n")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("[FinanceService] Failed to read message: %s\n", err.Error())
		}

		var event Event
		err = json.Unmarshal(msg.Value, &event)
		if err != nil {
			log.Fatalf("[FinanceService] Failed to parse JSON into event struct: %s\n", err.Error())
		}

		isSuccess := rand.Intn(2)

		if isSuccess == 0 {
			event.Status = "student.registration_failed"
		} else {
			event.Status = "student.registration_validated"
			log.Printf("[FinanceService] Payment validated for student_id: %s\n", event.StudentID)
		}

		writer := kafka.Writer{
			Addr:  kafka.TCP("localhost:9092"),
			Topic: event.Status,
		}

		eventByte, err := json.Marshal(event)
		if err != nil {
			log.Fatalf("[FinanceService] Failed to encode event to JSON: %s\n", err.Error())
		}

		log.Printf("[FinanceService] Sent event: %s\n", event.Status)

		err = writer.WriteMessages(
			context.TODO(),
			kafka.Message{
				Key:   []byte("Key-1"),
				Value: eventByte,
			},
		)
		if err != nil {
			log.Fatalf("[FinanceService] Failed to write message: %s\n", err)
		}
	}
}
