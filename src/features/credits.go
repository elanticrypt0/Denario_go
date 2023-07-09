package features

import (
	"strconv"

	"github.com/elanticrypt0/denario_go/src/models"
	"github.com/gofiber/fiber/v2"
)

func FindOneCredit(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id", "0"))
	return c.JSON(models.FindOneCredit(id))
}

func FindAllCredits(c *fiber.Ctx) error {
	return c.JSON(models.FindAllCredits())
}

func CreateCredit(c *fiber.Ctx) error {
	name := c.Params("name", "")
	comment := c.Params("comment", "")
	amount, _ := strconv.ParseFloat(c.Params("amount", "0"), 64)
	payments, _ := strconv.Atoi(c.Params("payments", "0"))
	started_at := c.Params("started_at", "")
	finished_at := c.Params("finished_at", "")
	category_id, _ := strconv.Atoi(c.Params("category_id", "0"))
	is_mutable, _ := strconv.ParseBool(c.Params("is_mutable", "false"))

	credit := models.CreateCredit(name, comment, amount, payments, started_at, finished_at, category_id, is_mutable)

	return c.JSON(credit)
}

func UpdateCredit(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id", "0"))
	credit := models.FindOneCredit(id)
	name := c.Params("name", "")
	comment := c.Params("comment", "")
	amount, _ := strconv.ParseFloat(c.Params("amount", "0"), 64)
	payments, _ := strconv.Atoi(c.Params("payments", "0"))
	started_at := c.Params("started_at", "")
	finished_at := c.Params("finished_at", "")

	credit.Name = name
	credit.Comment = comment
	credit.Amount = amount
	credit.Payments = payments
	credit.StartedAt = started_at
	credit.FinishedAt = finished_at

	credit = models.UpdateCredit(credit)

	return c.JSON(credit)
}

func DeleteCredit(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id", "0"))
	return c.JSON(models.DeleteCredit(id))
}
