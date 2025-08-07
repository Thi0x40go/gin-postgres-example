package controllers

import (
	"myapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetWidgets(c *gin.Context, db *gorm.DB) {
	var widgets []models.Widget
	if err := db.Find(&widgets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar widgets"})
		return
	}
	c.JSON(http.StatusOK, widgets)
}
