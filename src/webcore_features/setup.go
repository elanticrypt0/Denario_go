package webcore_features

import (
	"github.com/elanticrypt0/denario_go/src/models"
	"github.com/elanticrypt0/denario_go/src/webcore"
	"github.com/gofiber/fiber/v2"
)

func Setup(c *fiber.Ctx) error {
	migrateModels()
	return c.SendString("Setup enabled. Models Migrated.")
}

func migrateModels() {
	feature := webcore.NewFeature()
	feature.Db.AutoMigrate(&models.Category{}, &models.Credit{}, &models.Record{})
}
