package features

import (
	"strconv"

	"github.com/elanticrypt0/denario_go/src/models"
	"github.com/gofiber/fiber/v2"
)

// Find a record by ID
// Return the record
// Return an error if there was a problem
func FindOneRecord(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id", "0"))
	return c.JSON(models.FindOneRecord(id))
}

// Find all records
// Return the records
// Return an error if there was a problem
func FindAllRecords(c *fiber.Ctx) error {
	return c.JSON(models.FindAllRecords())
}

// Create a new record
// Return the new record
// Return an error if there was a problem
func CreateRecord(c *fiber.Ctx) error {
	r := new(models.Record)
	c.BodyParser(&r)
	record := models.CreateRecord(r.Name, r.Amount, r.AmountIo, r.Comment, r.RecordDate, r.CategoryID, r.IsMutable)
	return c.JSON(record)
}

// Update a record
// Return the updated record
// Return an error if there was a problem
func UpdateRecord(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id", "0"))
	record := models.FindOneRecord(id)
	c.BodyParser(&record)
	models.UpdateRecord(record)
	return c.JSON(record)
}

// Delete a record
// Return the deleted record
// Return an error if there was a problem
func DeleteRecord(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id", "0"))
	return c.JSON(models.DeleteRecord(id))
}
