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
	dbConnection := db.ConnectDB()
	defer func(dbConnection *sql.DB) {
		err := dbConnection.Close()
		if err != nil {
			log.Printf("Ocurrio un error: %v", err)
		}
	}(dbConnection)

	// ::: Initialize Gin router
	router := gin.Default()

	// ::: Routes
	router.GET("/api/novels", handlers.GetNovelsHandler(dbConnection))
	router.GET("/api/genres", handlers.GetGenresHandler(dbConnection))

	// ::: Start the Gin server
	router.Run(":8080")
}
