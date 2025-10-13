package main

import (
	"log"

	"github.com/carlospinellowork/drinks-online-fullstack/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	routes.SetupRoutes(app)

	log.Println("Servidor rodando em http://localhost:3333")
	app.Listen(":3333")
}
