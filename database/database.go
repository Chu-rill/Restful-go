package database

import (
	"log"
	"os"

	"github.com/Chu-rill/Restful-go/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct{
	Db *gorm.DB
}

var Databse Dbinstance

func ConnectDb(){
	err := godotenv.Load() // Load .env
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	dsn := os.Getenv("DATABASE_URL")

	db,err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		os.Exit(2)
	}

	log.Println("Connected to PostgreSQL successfully!")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations")

	// Auto Migrate
	db.AutoMigrate(&models.User{},&models.Product{},&models.Order{})

	Databse = Dbinstance{Db:db}
}