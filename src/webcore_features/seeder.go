package webcore_features

import (
	"encoding/json"
	"log"

	"github.com/elanticrypt0/denario_go/src/models"
	"github.com/elanticrypt0/denario_go/src/webcore/handlers"
	"github.com/gofiber/fiber/v2"
)

func Seed(c *fiber.Ctx) error {
	insertCategories()
	return c.JSON("OK")
}

func seedTable(table_name string) {

}

func insertCategories() {
	file := handlers.ReadJsonFile("categories")
	cat := []models.Category{}
	err := json.Unmarshal(file, &cat)
	if err != nil {
		log.Fatal("Cant unmarshal json", err)
	}
	for _, category := range cat {
		models.CreateCategory(category.Name)
	}
}
