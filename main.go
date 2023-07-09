package main

import (
	"fmt"

	"github.com/elanticrypt0/denario_go/src/features"
	"github.com/elanticrypt0/denario_go/src/webcore"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app_config := webcore.LoadConfig()
	fmt.Println(app_config.App_url)

	app := fiber.New()

	api := app.Group("/api")
	categories := api.Group("/categories")
	// categories
	categories.Get("/", func(c *fiber.Ctx) error {
		return features.FindAllCategories(c)
	})
	categories.Get("/:id", func(c *fiber.Ctx) error {
		return features.FindOneCategory(c)
	})

	// credits
	credits := api.Group("/credits")
	credits.Get("/", func(c *fiber.Ctx) error {
		return features.FindAllCredits(c)
	})

	// records
	records := api.Group("/records")
	records.Get("/", func(c *fiber.Ctx) error {
		return features.FindAllRecords(c)
	})

	// setup
	setup := app.Group("/setup")
	if app_config.App_setup_enabled {
		setup.Get("/setup", func(c *fiber.Ctx) error {
			return c.SendString("Setup enabled")
		})

		setup.Get("/seeder", func(c *fiber.Ctx) error {
			return c.SendString("Seeder enabled")
		})
	}

	app.Static("/", "./public")

	app.Listen(":" + app_config.App_server_port)

	fmt.Println("Server started on port 3000")

}
