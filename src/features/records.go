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
	name := c.Params("name", "")
	amount, _ := strconv.ParseFloat(c.Params("amount", "0"), 64)
	amount_io := c.Params("amount_io", "")
	comment := c.Params("comment", "")
	record_date := c.Params("record_date", "")
	category_id, _ := strconv.Atoi(c.Params("category_id", "0"))
	is_mutable, _ := strconv.ParseBool(c.Params("is_mutable", "false"))

	record := models.CreateRecord(name, amount, amount_io, comment, record_date, category_id, is_mutable)
	return c.JSON(record)
}

// Update a record
// Return the updated record
// Return an error if there was a problem
func UpdateRecord(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id", "0"))
	name := c.Params("name", "")
	amount, _ := strconv.ParseFloat(c.Params("amount", "0"), 64)
	amount_io := c.Params("amount_io", "")
	comment := c.Params("comment", "")
	record_date := c.Params("record_date", "")
	category_id, _ := strconv.Atoi(c.Params("category_id", "0"))
	is_mutable, _ := strconv.ParseBool(c.Params("is_mutable", "false"))

	record := models.FindOneRecord(id)
	record.Name = name
	record.Amount = amount
	record.AmountIo = amount_io
	record.Comment = comment
	record.RecordDate = record_date
	record.CategoryID = category_id
	record.IsMutable = is_mutable
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
