package handlers

import "github.com/gofiber/fiber/v2"

var mockDrinks = []map[string]interface{}{
	{"id": 1, "name": "Mojito", "category": "Cocktail", "price": 25.5},
	{"id": 2, "name": "Caipirinha", "category": "Cocktail", "price": 22.0},
	{"id": 3, "name": "Gin Tônica", "category": "Gin", "price": 28.9},
}

func GetDrinks(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"data": mockDrinks,
	})
}
