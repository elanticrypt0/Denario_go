package webcore

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
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
		Posgres_url:       os.Getenv("POSTGRES_URL"),
	}

}

func LoadConfigAndConnectDB() (*gorm.DB, AppConfig) {
	app_config := LoadConfig()
	db_connection := DbConnect(app_config.Posgres_url)
	return db_connection, app_config
}

func NewFeature() *Features {
	db, app_config := LoadConfigAndConnectDB()
	return &Features{
		Db:        db,
		AppConfig: app_config,
	}
}
