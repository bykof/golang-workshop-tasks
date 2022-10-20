package main

import (
	"database_example/core/ports"
	"database_example/infrastructure/database/repositories"
	"database_example/interfaces/controllers"
)

func main() {
	var err error
	var productRepository ports.ProductRepository

	productRepository = repositories.NewSqLiteProductRepository()
	productController := controllers.NewProductController(productRepository)

}
