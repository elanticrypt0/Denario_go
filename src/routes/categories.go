package routes

// func Categories(api *fiber.Router) *fiber.Router {
// 	// categories
// 	categories := api.Group("/categories")
// 	categories.Get("/", func(c *fiber.Ctx) error {
// 		return features.FindAllCategories(c)
// 	})
// 	categories.Get("/:id", func(c *fiber.Ctx) error {
// 		return features.FindOneCategory(c)
// 	})

// 	categories.Post("/", func(c *fiber.Ctx) error {
// 		return features.CreateCategory(c)
// 	})

// 	categories.Put("/:id", func(c *fiber.Ctx) error {
// 		return features.UpdateCategory(c)
// 	})

// 	categories.Delete("/:id", func(c *fiber.Ctx) error {
// 		return features.DeleteCategory(c)
// 	})
// 	return api
// }
