package handlers

import (
	"strconv"
	"strings"

	"github.com/carlospinellowork/drinks-online-fullstack/internal/database"
	"github.com/carlospinellowork/drinks-online-fullstack/internal/models"
	"github.com/gofiber/fiber/v2"
)

func GetDrinks(c *fiber.Ctx) error {
	var drinks []models.Drink

	category := c.Query("category")
	search := c.Query("search")
	pageStr := c.Query("page", "1")
	limitStr := c.Query("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	offset := (page - 1) * limit

	query := database.DB.Model(&models.Drink{})

	if category != "" {
		query = query.Where("LOWER(category) = ?", strings.ToLower(category))
	}

	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}

	result := query.Limit(limit).Offset(offset).Find(&drinks)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao buscar drinks"})
	}

	return c.JSON(fiber.Map{"data": drinks})
}

func CreateDrink(c *fiber.Ctx) error {
	drink := new(models.Drink)

	if err := c.BodyParser(drink); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Dados inválidos"})
	}

	result := database.DB.Create(&drink)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao criar drink"})
	}

	return c.Status(201).JSON(fiber.Map{"data": drink})
}

func UpdateDrink(c *fiber.Ctx) error {
	id := c.Params("id")
	var drink models.Drink

	if err := database.DB.First(&drink, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Drink não encontrado"})
	}

	if err := c.BodyParser(&drink); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Dados inválidos"})
	}

	database.DB.Save(&drink)

	return c.JSON(fiber.Map{"data": drink})
}
