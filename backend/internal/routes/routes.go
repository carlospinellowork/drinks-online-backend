package routes

import (
	"github.com/carlospinellowork/drinks-online-fullstack/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/")

	api.Get("/drinks", handlers.GetDrinks)
}
