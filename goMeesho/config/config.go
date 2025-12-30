package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetMongoURI() string {
	return os.Getenv("MONGO_URI")
}

func GetMongoDBName() string {
	return os.Getenv("MONGO_DB_NAME")
}

func GetKafkaServer() string {
	return os.Getenv("KAFKA_SERVER")
}

func GetKafkaTopic() string {
	return os.Getenv("KAFKA_TOPIC_NAME")
}

func GetKafkaGroupID() string {
	return os.Getenv("KAFKA_GROUP_ID")
}
