package routes

import (
	"github.com/carlospinellowork/drinks-online-fullstack/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/drinks")

	api.Get("/", handlers.GetDrinks)
	api.Post("/", handlers.CreateDrink)
	api.Patch("/:id", handlers.UpdateDrink)
}
