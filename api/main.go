package main

import (
	"fmt"

	"github.com/elanticrypt0/denario_go/core"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app_config := core.LoadConfig()
	// app_url := fmt.Sprintf("%s:%s", appConfig.App_server_host, appConfig.App_server_port)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(app_config.App_server_port)

	fmt.Println("Server started on port 3000")

}
