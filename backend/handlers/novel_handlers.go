package handlers

import (
	"backend/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetNovelsHandler(database *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		novels, err := db.GetAllNovels(database)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"titles": novels})
	}
}

func GetGenresHandler(database *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		genres, err := db.GetAllGenres(database)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"genres": genres})
	}
}
