package main

import (
	"backend/db"
	"backend/handlers"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	// ::: Connection to the database
	database := db.ConnectDB()
	defer func(database *sql.DB) {
		err := database.Close()
		if err != nil {
			log.Printf("Ocurrio un error: %v", err)
		}
	}(database)

	// ::: Initialize Gin router
	router := gin.Default()

	// ::: Get all novels
	router.GET("/api/novels", handlers.GetNovelsHandler(database))

	// ::: Start the Gin server
	router.Run(":8080")
}
