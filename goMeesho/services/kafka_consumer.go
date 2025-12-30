package services

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"goMeesho/config"
	"goMeesho/models"

	"github.com/segmentio/kafka-go"
)

func StartKafkaConsumer() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{config.GetKafkaServer()},
		Topic:   config.GetKafkaTopic(),
		GroupID: config.GetKafkaGroupID(),
	})

	defer reader.Close()

	log.Println("Kafka consumer started")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("Kafka read error:", err)
		}

		var incoming models.IncomingSMS
		if err := json.Unmarshal(msg.Value, &incoming); err != nil {
			log.Println("Invalid Kafka message:", err)
			continue
		}

		sms := models.SMS{
			PhoneNumber: incoming.PhoneNumber,
			Messages: []models.Message{
				{
					Text:      incoming.Message,
					Status:    incoming.Status,
					CreatedAt: time.Now(),
				},
			},
		}

		log.Println("Received Kafka message:", sms)

		if err := SaveSMS(context.Background(), &sms); err != nil {
			log.Println("Mongo upsert failed:", err)
		}
	}
}
