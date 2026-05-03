package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// AppConfig holds all application configuration values.
type AppConfig struct {
	DBURL         string
	JWTSecret     string
	JWTExpiryHours int
}

var Env AppConfig

// LoadEnv reads values from .env file and populates the AppConfig.
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, falling back to system environment variables")
	}

	Env = AppConfig{
		DBURL:         getEnv("DB_URL", "host=localhost user=postgres password=P@ssw0rd dbname=gin_db1 port=5432 sslmode=disable"),
		JWTSecret:     getEnv("JWT_SECRET", "default-secret-change-me"),
		JWTExpiryHours: getEnvAsInt("JWT_EXPIRY_HOURS", 24),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if value, exists := os.LookupEnv(key); exists {
		i, err := strconv.Atoi(value)
		if err == nil {
			return i
		}
	}
	return fallback
}
