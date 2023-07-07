package core

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	App_server_host   string
	App_server_port   string
	App_setup_enabled bool
	App_debug_mode    bool
	Posgres_url       string
}

// Load the .env vars and return a struct
func LoadConfig() AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return AppConfig{
		App_server_host:   os.Getenv("APP_SERVER_HOST"),
		App_server_port:   os.ExpandEnv("APP_SERVER_PORT"),
		App_setup_enabled: os.ExpandEnv("APP_SETUP_ENABLED") == "true",
		App_debug_mode:    os.ExpandEnv("APP_DEBUG_MODE") == "true",
		Posgres_url:       "postgres://postgres:postgres@localhost:5432/denario?sslmode=disable",
	}
}
