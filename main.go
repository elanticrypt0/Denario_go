package main

import (
	"fmt"

	"github.com/elanticrypt0/denario_go/features"
	"github.com/elanticrypt0/denario_go/webcore"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app_config := webcore.LoadConfig()
	fmt.Println(app_config.App_url)

	app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World 👋!")
	// })

	// categories
	app.Get("/categories", func(c *fiber.Ctx) error {
		categories := features.FindAllCategories()
		return c.JSON(categories)
	})

	// credits
	app.Get("/credits", func(c *fiber.Ctx) error {
		credits := features.FindAllCredits()
		return c.JSON(credits)
	})

	if app_config.App_setup_enabled {
		app.Get("/setup", func(c *fiber.Ctx) error {
			return c.SendString("Setup enabled")
		})
	}

	app.Static("/", "./public")

	app.Listen(":" + app_config.App_server_port)

	fmt.Println("Server started on port 3000")

}
