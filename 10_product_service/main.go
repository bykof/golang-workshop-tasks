package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	productController, err := InitializeProductController()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.GET("/products", productController.List)
	r.GET("/products/:id", productController.Get)
	r.PUT("/products/:id", productController.Update)
	r.DELETE("/products/:id", productController.Delete)
	r.POST("/products", productController.Add)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
