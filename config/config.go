package config

import (
	"net/http"
	"os"
)

type AppConfig struct {
	HttpClient *http.Client
	BaseURL    string
	Port       string
}

func NewAppConfig(client *http.Client) *AppConfig {
	cfg := &AppConfig{}

	cfg.HttpClient = client
	cfg.BaseURL = getEnv("WEATHER_API_URL", "https://api.open-meteo.com/v1/forecast")
	cfg.Port = getEnv("PORT", "8080")

	return cfg
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
