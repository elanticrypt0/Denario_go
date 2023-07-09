package webcore

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnect(dns string) *gorm.DB {
	// connect to gorn
	conn, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	return conn
}
