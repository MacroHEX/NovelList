package db

import (
	"backend/models"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func ConnectDB() *sql.DB {

	// ::: Read from .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Printf("Ocurrio un error: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	// ::: Connect to the database
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
		dbHost, dbPort, dbUser, dbPassword, dbName))

	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	return db
}

func GetAllNovels(db *sql.DB) ([]models.LightNovel, error) {
	var novels []models.LightNovel
	rows, err := db.Query("SELECT LightNovelID, Title, NumberOfVolumes, CoverArtURL FROM LightNovels ORDER BY Title")

	if err != nil {
		log.Printf("Error fetching light novels: %v", err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf("Ocurrio un error: %v", err)
		}
	}(rows)

	for rows.Next() {
		var novel models.LightNovel
		if err := rows.Scan(&novel.LightNovelID, &novel.Title, &novel.NumberOfVolumes, &novel.CoverArtURL); err != nil {
			log.Printf("Error scanning novel: %v", err)
			return nil, err
		}
		novels = append(novels, novel)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error con columnas: %v", err)
	}

	return novels, nil
}
