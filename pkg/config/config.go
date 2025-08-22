package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv  string
	AppPort string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  No .env file found, using system environment variables")
	}

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		AppEnv:  env,
		AppPort: port,
	}
}
