package webcore

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

func DbConnect(host, user, password, dbname, port string) *gorm.DB {

	// dns := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Argentina/Buenos_Aires"
	const dns = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable "
	// connect to gorn
	conn, err := gorm.Open(postgres.Open(fmt.Sprintf(dns, host, user, password, dbname, port)), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	return conn
}
