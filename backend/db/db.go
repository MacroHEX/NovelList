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

	// ::: Connect to the database
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, sslMode))

	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	return db
}

func GetAllNovels(db *sql.DB) ([]models.LightNovel, error) {
	var novels []models.LightNovel

	// Initial query to fetch novels
	rows, err := db.Query(`
        SELECT LightNovelID, Title, NumberOfVolumes, CoverArtURL
        FROM LightNovels
        ORDER BY Title
    `)
	if err != nil {
		log.Printf("Error fetching novels: %v", err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		var novel models.LightNovel
		err := rows.Scan(&novel.LightNovelID, &novel.Title, &novel.NumberOfVolumes, &novel.CoverArtURL)
		if err != nil {
			log.Printf("Error scanning novel: %v", err)
			continue
		}

		novel.Volumes = fetchVolumesForNovel(db, novel.LightNovelID)
		novel.Genres = fetchGenresForNovel(db, novel.LightNovelID)
		novel.Authors = fetchAuthorsForNovel(db, novel.LightNovelID)

		novels = append(novels, novel)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error with row processing: %v", err)
		return nil, err
	}

	return novels, nil
}

func GetAllGenres(db *sql.DB) ([]models.Genre, error) {
	var genres []models.Genre
	rows, err := db.Query("SELECT GenreID, Name FROM Genres ORDER BY Name")

	if err != nil {
		log.Printf("Error fetching genres: %v", err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf("Ocurrio un error: %v", err)
		}
	}(rows)

	for rows.Next() {
		var genre models.Genre
		if err := rows.Scan(&genre.GenreID, &genre.Name); err != nil {
			log.Printf("Error scanning novel: %v", err)
			return nil, err
		}
		genres = append(genres, genre)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error con las columnas: %v", err)
	}

	return genres, nil
}

// ::: Internal
func fetchVolumesForNovel(db *sql.DB, novelID int) []models.Volume {
	var volumes []models.Volume
	volumeRows, err := db.Query(`
        SELECT VolumeID, LightNovelID, VolumeNumber, CoverArtURL, ReleaseDate
        FROM Volumes
        WHERE LightNovelID = $1
        ORDER BY VolumeNumber
    `, novelID)
	if err != nil {
		log.Printf("Error fetching volumes for novel %d: %v", novelID, err)
		return volumes
	}
	defer func(volumeRows *sql.Rows) {
		err := volumeRows.Close()
		if err != nil {

		}
	}(volumeRows)

	for volumeRows.Next() {
		var volume models.Volume
		if err := volumeRows.Scan(&volume.VolumeID, &volume.LightNovelID, &volume.VolumeNumber, &volume.CoverArtURL, &volume.ReleaseDate); err != nil {
			log.Printf("Error scanning volume: %v", err)
			continue
		}
		volumes = append(volumes, volume)
	}
	return volumes
}

func fetchGenresForNovel(db *sql.DB, novelID int) []models.Genre {
	var genres []models.Genre
	genreRows, err := db.Query(`
        SELECT g.GenreID, g.Name
        FROM Genres g
        JOIN LightNovelGenres lng ON g.GenreID = lng.GenreID
        WHERE lng.LightNovelID = $1
        ORDER BY g.Name
    `, novelID)
	if err != nil {
		log.Printf("Error fetching genres for novel %d: %v", novelID, err)
		return genres
	}
	defer func(genreRows *sql.Rows) {
		err := genreRows.Close()
		if err != nil {

		}
	}(genreRows)

	for genreRows.Next() {
		var genre models.Genre
		if err := genreRows.Scan(&genre.GenreID, &genre.Name); err != nil {
			log.Printf("Error scanning genre: %v", err)
			continue
		}
		genres = append(genres, genre)
	}
	return genres
}

func fetchAuthorsForNovel(db *sql.DB, novelID int) []models.Author {
	var authors []models.Author
	authorRows, err := db.Query(`
        SELECT a.AuthorID, a.Name
        FROM Authors a
        JOIN LightNovelAuthors lna ON a.AuthorID = lna.AuthorID
        WHERE lna.LightNovelID = $1
        ORDER BY a.Name
    `, novelID)
	if err != nil {
		log.Printf("Error fetching authors for novel %d: %v", novelID, err)
		return authors
	}
	defer func(authorRows *sql.Rows) {
		err := authorRows.Close()
		if err != nil {

		}
	}(authorRows)

	for authorRows.Next() {
		var author models.Author
		if err := authorRows.Scan(&author.AuthorID, &author.Name); err != nil {
			log.Printf("Error scanning author: %v", err)
			continue
		}
		authors = append(authors, author)
	}
	return authors
}
