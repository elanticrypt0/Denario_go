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
	App_url           string
	App_setup_enabled bool
	App_debug_mode    bool
	Db_config         DatabaseConfig
}

// Load the .env vars and return a struct
func LoadConfig() AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return AppConfig{
		App_server_host:   os.Getenv("APP_SERVER_HOST"),
		App_server_port:   os.Getenv("APP_SERVER_PORT"),
		App_url:           os.Getenv("APP_SERVER_HOST") + ":" + os.Getenv("APP_SERVER_PORT"),
		App_setup_enabled: os.Getenv("APP_SETUP_ENABLED") == "true",
		App_debug_mode:    os.Getenv("APP_DEBUG_MODE") == "true",
		Db_config: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Dbname:   os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
		},
	}

}

func LoadConfigAndConnectDB() (*gorm.DB, AppConfig) {
	app_config := LoadConfig()
	db_connection := DbConnect(app_config.Db_config.Host, app_config.Db_config.User, app_config.Db_config.Password, app_config.Db_config.Dbname, app_config.Db_config.Port)
	return db_connection, app_config
}

func NewFeature() *Features {
	db, app_config := LoadConfigAndConnectDB()
	return &Features{
		Db:        db,
		AppConfig: app_config,
	}
}
