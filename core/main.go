package core

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	app_server_host   string
	app_server_port   string
	app_setup_enabled bool
	app_debug_mode    bool
	posgres_url       string
}

// Load the .env vars and return a struct
func LoadConfig() AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// put everything in a struct
	// return the struct
	return AppConfig{
		app_server_host:   os.Getenv("APP_SERVER_HOST"),
		app_server_port:   os.ExpandEnv("APP_SERVER_PORT"),
		app_setup_enabled: os.ExpandEnv("APP_SETUP_ENABLED") == "true",
		app_debug_mode:    os.ExpandEnv("APP_DEBUG_MODE") == "true",
		posgres_url:       "postgres://postgres:postgres@localhost:5432/denario?sslmode=disable",
	}
}
