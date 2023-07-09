package webcore_features

import "github.com/gofiber/fiber/v2"

func Status(c *fiber.Ctx) error {
	return c.JSON("OK")
}

// ReadEnv reads the environment
// Return the names and values of the environment variables
func ReadEnv(c *fiber.Ctx) error {
	return c.JSON("OK")
}
