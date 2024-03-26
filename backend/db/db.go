package db

import (
	"backend/models"
	"fmt"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "local"
	}

	envFileName := fmt.Sprintf(".env.%s", appEnv)

	// ::: Read from .env file
	if err := godotenv.Load(envFileName); err != nil {
		log.Fatalf("Error loading %s file", envFileName)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("SSL")

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPassword, dbName, sslMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	return db
}

func GetAllNovels(db *gorm.DB) ([]models.LightNovel, error) {
	var novels []models.LightNovel
	result := db.Preload("Volumes").Preload("Genres").Preload("Authors").Order("title").Find(&novels)
	if result.Error != nil {
		log.Printf("Error fetching novels: %v", result.Error)
		return nil, result.Error
	}

	return novels, nil
}

func GetAllGenres(db *gorm.DB) ([]models.Genre, error) {
	var genres []models.Genre
	result := db.Order("name").Find(&genres)
	if result.Error != nil {
		log.Printf("Error fetching genres: %v", result.Error)
		return nil, result.Error
	}

	return genres, nil
}
