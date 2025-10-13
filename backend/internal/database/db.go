package database

import (
	"fmt"
	"log"

	"github.com/carlospinellowork/drinks-online-fullstack/internal/config"
	"github.com/carlospinellowork/drinks-online-fullstack/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	config.LoadEnv()

	host := config.GetEnv("DB_HOST")
	port := config.GetEnv("DB_PORT")
	user := config.GetEnv("DB_USER")
	password := config.GetEnv("DB_PASSWORD")
	dbname := config.GetEnv("DB_NAME")
	sslmode := config.GetEnv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}

	err = db.AutoMigrate(&models.Drink{})

	if err != nil {
		log.Fatal("Erro ao migrar o banco:", err)
	}

	DB = db
	log.Println("Conectado ao banco de dados!")
}
