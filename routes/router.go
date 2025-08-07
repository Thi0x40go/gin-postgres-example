package routes

import (
	"myapp/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/widgets", func(c *gin.Context) {
		controllers.GetWidgets(c, db)
	})

	return r
}
