package models

import (
	"gorm.io/gorm"
)

type Drink struct {
	gorm.Model
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	ImageURL string  `json:"image_url"`
}
