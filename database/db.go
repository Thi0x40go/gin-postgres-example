package database

import (
	"log"
	"myapp/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}

	// AutoMigrate cria as tabelas se necessário
	err = db.AutoMigrate(&models.Widget{})
	if err != nil {
		log.Fatalf("Erro no AutoMigrate: %v", err)
	}

	DB = db
	return db
}
