package main

import (
	"fmt"

	"github.com/elanticrypt0/denario_go/src/features"
	"github.com/elanticrypt0/denario_go/src/webcore"
	"github.com/elanticrypt0/denario_go/src/webcore_features"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app_config := webcore.LoadConfig()
	fmt.Println(app_config.App_url)

	app := fiber.New()

	api := app.Group("/api")

	// categories
	categories := api.Group("/categories")
	categories.Get("/", func(c *fiber.Ctx) error {
		return features.FindAllCategories(c)
	})
	categories.Get("/:id", func(c *fiber.Ctx) error {
		return features.FindOneCategory(c)
	})

	categories.Post("/", func(c *fiber.Ctx) error {
		return features.CreateCategory(c)
	})

	categories.Put("/:id", func(c *fiber.Ctx) error {
		return features.UpdateCategory(c)
	})

	categories.Delete("/:id", func(c *fiber.Ctx) error {
		return features.DeleteCategory(c)
	})

	// credits
	credits := api.Group("/credits")
	credits.Get("/", func(c *fiber.Ctx) error {
		return features.FindAllCredits(c)
	})

	credits.Get("/:id", func(c *fiber.Ctx) error {
		return features.FindOneCredit(c)
	})

	credits.Post("/:category_id", func(c *fiber.Ctx) error {
		return features.CreateCredit(c)
	})

	credits.Put("/:id", func(c *fiber.Ctx) error {
		return features.UpdateCredit(c)
	})

	credits.Delete("/:id", func(c *fiber.Ctx) error {
		return features.DeleteCredit(c)
	})

	// records
	records := api.Group("/records")
	records.Get("/", func(c *fiber.Ctx) error {
		return features.FindAllRecords(c)
	})

	records.Get("/:id", func(c *fiber.Ctx) error {
		return features.FindOneRecord(c)
	})

	records.Post("/", func(c *fiber.Ctx) error {
		return features.CreateRecord(c)
	})

	records.Put("/:id", func(c *fiber.Ctx) error {
		return features.UpdateRecord(c)
	})

	records.Delete("/:id", func(c *fiber.Ctx) error {
		return features.DeleteRecord(c)
	})

	// setup
	setup := api.Group("/setup")
	if app_config.App_setup_enabled {
		setup.Get("/", func(c *fiber.Ctx) error {
			return webcore_features.Setup(c)
		})
	}

	//status
	status := api.Group("/status")
	status.Get("/", func(c *fiber.Ctx) error {
		return webcore_features.Status(c)
	})

	status.Get("/getenv", func(c *fiber.Ctx) error {
		return webcore_features.ReadEnv(c)
	})

	seeder := api.Group("/seeder")
	if app_config.App_setup_enabled {
		seeder.Get("/seed/", func(c *fiber.Ctx) error {
			return webcore_features.Seed(c)
		})
		seeder.Get("/seed/:table_name", func(c *fiber.Ctx) error {
			return webcore_features.Seed(c)
		})
	}

	app.Static("/", "./public")

	app.Listen(":" + app_config.App_server_port)

	fmt.Println("Server started on port 3000")

}
