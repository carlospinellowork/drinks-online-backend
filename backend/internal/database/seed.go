package database

import (
	"log"

	"github.com/carlospinellowork/drinks-online-fullstack/internal/models"
)

func Seed() {
	var count int64
	DB.Model(&models.Drink{}).Count(&count)

	if count > 0 {
		return
	}

	drinks := []models.Drink{
		// Categoria: DRINKS
		{
			Name:     "Caipirinha de Limão",
			Category: "drinks",
			Price:    22.00,
			ImageURL: "https://images.unsplash.com/photo-1513558161293-cdaf765ed2fd?q=80&w=800",
		},
		{
			Name:     "Moscow Mule",
			Category: "drinks",
			Price:    32.00,
			ImageURL: "https://images.unsplash.com/photo-1513558161293-cdaf765ed2fd?q=80&w=800",
		},
		{
			Name:     "Gin Tônica Clássico",
			Category: "drinks",
			Price:    28.50,
			ImageURL: "https://images.unsplash.com/photo-1514362545857-3bc16c4c7d1b?q=80&w=800",
		},

		// Categoria: WHISKEYS (Bebidas)
		{
			Name:     "Jack Daniel's Old No. 7",
			Category: "whiskeys",
			Price:    150.00,
			ImageURL: "https://images.unsplash.com/photo-1527281480448-f357169ab5e9?q=80&w=800",
		},
		{
			Name:     "Johnnie Walker Black Label",
			Category: "whiskeys",
			Price:    180.00,
			ImageURL: "https://images.unsplash.com/photo-1527281480448-f357169ab5e9?q=80&w=800",
		},

		// Categoria: BEERS (Cervejas)
		{
			Name:     "Heineken 330ml",
			Category: "beers",
			Price:    9.50,
			ImageURL: "https://images.unsplash.com/photo-1538944513120-5c66046da896?q=80&w=800",
		},
		{
			Name:     "Corona Extra",
			Category: "beers",
			Price:    11.00,
			ImageURL: "https://images.unsplash.com/photo-1625640244580-049830574176?q=80&w=800",
		},
		{
			Name:     "Stella Artois",
			Category: "beers",
			Price:    10.00,
			ImageURL: "https://images.unsplash.com/photo-1597403294029-fcc6191144f2?q=80&w=800",
		},
	}

	for _, drink := range drinks {
		DB.Create(&drink)
	}

	log.Println("Banco de dados populado com sucesso!")
}
