package routes

import (
	"github.com/efectn/rest-countries/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterAPIRoutes(route fiber.Router) {
	route = route.Group("/api/v1")

	route.Get("/all", controllers.GetAll)
	route.Get("/name/:name", controllers.GetByName)
	route.Get("/code/:code", controllers.GetByCode)
	route.Get("/currency/:currency", controllers.GetByCurrrency)
	route.Get("/language/:language", controllers.GetByLanguage)
	route.Get("/capital/:capital", controllers.GetByCapital)
	route.Get("/calling/:code", controllers.GetByCallingCode)
	route.Get("/region/:region", controllers.GetByRegion)
	route.Get("/regional-block/:block", controllers.GetByRegionalBlock)

	route.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Endpoint not found!",
		})
	})
}
