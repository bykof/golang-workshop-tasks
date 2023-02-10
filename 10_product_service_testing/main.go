package main

import (
	"database_example/infrastructure/database/orm_entities"
	"database_example/infrastructure/database/repositories"
	"database_example/interfaces/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRouter() *gin.Engine {
	var err error

	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&orm_entities.ORMProduct{})
	productRepository := repositories.NewSqLiteProductRepository(db)
	productController := controllers.NewProductController(productRepository)
	r := gin.Default()
	r.GET("/products", productController.List)
	r.POST("/products", productController.Add)
	r.DELETE("/products/:id", productController.Remove)
	r.PATCH("/products/:id", productController.Update)
	return r
}

func main() {
	r := setupRouter()
	r.Run()
}
