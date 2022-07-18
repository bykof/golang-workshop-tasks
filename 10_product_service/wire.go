//go:build wireinject
// +build wireinject

package main

import (
	"os"
	"product_service/core/ports"
	"product_service/infrastructure/repositories"
	"product_service/interfaces/controllers"

	"github.com/google/wire"
)

func NewProductRepository() (ports.ProductRepositoryPort, error) {
	switch os.Getenv("REPOSITORY") {
	case "inmemory":
		return repositories.NewInMemoryProductRepository(), nil
	case "jsonfile":
		return repositories.NewJSONFileProductRepository("/Users/michaelbykovski/workspace/daimler/golang_workshop/product_service/products.json")
	default:
		return repositories.NewInMemoryProductRepository(), nil
	}
}

func InitializeProductController() (*controllers.ProductController, error) {
	wire.Build(NewProductRepository, controllers.NewProductController)
	return &controllers.ProductController{}, nil
}
