package controllers

import (
	"application/package/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetAllUsers(db *gorm.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		var users []models.User

		data := db.Find(users)

		if data.Error != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": data.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, users)
	}
}
