// cmd/bot/config.go
package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	TelegramToken   string
	Debug           bool
	HealthPort      int
	EmergencyAPIURL string
}

func LoadConfig() AppConfig {
	_ = godotenv.Load(".env")

	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_TOKEN environment variable is required")
	}

	debug := false
	if v := os.Getenv("DEBUG"); v == "1" || v == "true" || v == "TRUE" {
		debug = true
	}

	healthPort := 8080
	if v := os.Getenv("HEALTH_PORT"); v != "" {
		if p, err := strconv.Atoi(v); err == nil {
			healthPort = p
		} else {
			log.Printf("Invalid HEALTH_PORT=%q, using default %d", v, healthPort)
		}
	}

	emergencyURL := os.Getenv("EMERGENCY_API_URL")

	log.Printf("Config loaded: debug=%v, health_port=%d, emergency_api_url=%q",
		debug, healthPort, emergencyURL)

	return AppConfig{
		TelegramToken:   token,
		Debug:           debug,
		HealthPort:      healthPort,
		EmergencyAPIURL: emergencyURL,
	}
}
