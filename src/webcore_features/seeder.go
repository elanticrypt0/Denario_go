package webcore_features

import "github.com/gofiber/fiber/v2"

func Seed(c *fiber.Ctx) error {
	return c.JSON("OK")
}

func seedTable(table_name string) {

}

// this function open a .json file and return the content into Object
func readJsonFile(file_name string) {

}
