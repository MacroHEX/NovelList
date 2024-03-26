package main

import (
	"backend/db"
	"backend/handlers"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// ::: Connection to the database
	dbConnection := db.ConnectDB()

	// ::: Initialize Gin router
	router := gin.Default()

	// ::: Routes
	router.GET("/api/novels", handlers.GetNovelsHandler(dbConnection))
	router.GET("/api/genres", handlers.GetGenresHandler(dbConnection))

	// ::: Start the Gin server
	router.Run(":8080")
}
