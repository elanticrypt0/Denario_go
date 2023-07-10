package webcore_features

import (
	"encoding/json"
	"log"

	"github.com/elanticrypt0/denario_go/src/models"
	"github.com/elanticrypt0/denario_go/src/webcore/handlers"
	"github.com/gofiber/fiber/v2"
)

func Seed(c *fiber.Ctx) error {
	// seedCategories()
	// seedCredits()
	seedRecords()
	return c.JSON("OK")
}

func seedTable(table_name string) {

}

func seedCategories() {
	file := handlers.ReadJsonFile("categories")
	cat := []models.Category{}
	err := json.Unmarshal(file, &cat)
	if err != nil {
		log.Fatal("Cant unmarshal json", err)
	}
	for _, category := range cat {
		models.CreateCategory(category.Name)
	}
	log.Println("Categories seeded")
}

func seedCredits() {
	file := handlers.ReadJsonFile("credits")
	toInsert := []models.Credit{}
	err := json.Unmarshal(file, &toInsert)
	if err != nil {
		log.Fatal("Cant unmarshal json", err)
	}
	for _, item := range toInsert {
		models.CreateCredit(item.Name, item.Comment, item.Amount, item.Payments, item.StartedAt, item.FinishedAt, item.CategoryID)
	}
	log.Println("Credits seeded")
}

func seedRecords() {
	file := handlers.ReadJsonFile("records")
	toInsert := []models.Record{}
	err := json.Unmarshal(file, &toInsert)
	if err != nil {
		log.Fatal("Cant unmarshal json", err)
	}
	for _, item := range toInsert {
		models.CreateRecord(item.Name, item.Amount, item.AmountIo, item.Comment, item.RecordDate, item.CategoryID, item.IsMutable)
	}
	log.Println("Records seeded")
}
