package features

import (
	"strconv"

	"github.com/elanticrypt0/denario_go/src/models"
	"github.com/gofiber/fiber/v2"
)

func FindOneCategory(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id", "0"))
	return c.JSON(models.FindOneCategory(id))
}

func FindAllCategories(c *fiber.Ctx) error {
	categories := models.FindAllCategories()
	return c.JSON(categories)
}

func CreateCategory(c *fiber.Ctx) error {
	name := c.Params("name", "")
	category := models.CreateCategory(name)
	return c.JSON(category)
}

func UpdateCategory(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id", "0"))
	category := models.FindOneCategory(id)
	name := c.Params("name", "")
	category.Name = name
	category = models.UpdateCategory(category)
	return c.JSON(category)
}

func DeleteCategory(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id", "0"))
	category := models.DeleteCategory(id)
	return c.JSON(category)
}
