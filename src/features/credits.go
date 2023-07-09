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

	new_c := new(models.Credit)
	c.BodyParser(&new_c)
	credit := models.CreateCredit(new_c.Name, new_c.Comment, new_c.Amount, new_c.Payments, new_c.StartedAt, new_c.FinishedAt, new_c.CategoryID)

	cat_id := c.Params("category_id", "0")
	category_id, _ := strconv.Atoi(cat_id)
	models.CreateRecordsForCredit(credit, category_id)
	return c.JSON(credit)
}

func UpdateCredit(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id", "0"))
	credit := models.FindOneCredit(id)
	c.BodyParser(&credit)
	credit = models.UpdateCredit(credit)

	return c.JSON(credit)
}

func DeleteCredit(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id", "0"))
	return c.JSON(models.DeleteCredit(id))
}
