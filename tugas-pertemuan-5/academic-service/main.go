package main

import (
	"context"
	"encoding/json"
	"log"

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
		Topic:   "student.registration_validated",
		GroupID: "academic-service",
	})
	defer reader.Close()

	log.Println("[AcademicService] Started")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("[AcademicService] Failed to read message: %s\n", err.Error())
		}

		var event Event
		err = json.Unmarshal(msg.Value, &event)
		if err != nil {
			log.Printf("[AcademicService] Failed to parse JSON into event struct: %s\n", err.Error())
		}

		log.Printf("[AcademicService] Received event: %s, student_id: %s\n", event.Status, event.StudentID)

		event.Status = "student.academic_initialized"

		log.Printf("[AcademicService] Academic initialized for student_id: %s\n", event.StudentID)
	}
}
