package config

import (
	"fmt"
	"log"
	"os"

	"github.com/farhapartex/real_estate_be/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading env file")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if dbErr != nil {
		log.Fatal("Error connecting DB")
	}

	DB = db
	fmt.Println("DB connection successfull!")
}

func MigrateDB() {
	fmt.Println("Running DB migration ...")

	dbModels := []interface{}{
		&models.User{},
		&models.OwnerProfile{},
		&models.Country{},
		&models.Division{},
		&models.District{},
	}

	for _, model := range dbModels {
		err := DB.AutoMigrate(model)
		if err != nil {
			fmt.Printf("Error migrating %T: %v\n", model, err)
		}
	}

	fmt.Println("... DB migration completed.")
}
