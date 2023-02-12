package routes

import (
	"github.com/rizalnurizalludin/api-product/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/", controllers.Home)
	r.GET("/api/product", controllers.Index)
	r.POST("/api/product", controllers.Create)
	r.GET("/api/product/:id", controllers.Show)
	r.PATCH("/api/product/:id", controllers.Update)
	r.DELETE("api/product/:id", controllers.Delete)
	return r
}
