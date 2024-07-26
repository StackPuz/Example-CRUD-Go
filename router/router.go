package router

import (
	"app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() {
	productController := controllers.ProductController{}
	router := gin.Default()
	router.StaticFile("/", "./public/index.html")
	router.Group("/api").
		GET("/products", productController.Index).
		POST("/products", productController.Create).
		GET("/products/:id", productController.Get).
		PUT("/products/:id", productController.Update).
		DELETE("/products/:id", productController.Delete)
	router.Run()
}
