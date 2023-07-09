package webcore

import (
	"gorm.io/gorm"
)

type Features struct {
	Db        *gorm.DB
	AppConfig AppConfig
}
