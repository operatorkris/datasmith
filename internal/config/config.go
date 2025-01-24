package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	KafkaTopic string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default values.")
	}
	return &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		KafkaTopic: getEnv("KAFKA_TOPIC", "datasmith-events"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
