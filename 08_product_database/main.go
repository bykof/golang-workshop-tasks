package main

import (
	"database_example/core/domain/entities"
	"database_example/core/ports"
	"database_example/infrastructure/database/repositories"
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	var err error
	var productRepository ports.ProductRepository

	productRepository = repositories.NewSqLiteProductRepository()
	mentosMint := entities.Product{
		Name:  "Mentos Mint",
		Price: decimal.NewFromFloat(1.99),
	}
	mentosFruit := entities.Product{
		Name:  "Mentos Fruit",
		Price: decimal.NewFromFloat(1.49),
	}
	err = productRepository.Add(&mentosFruit)
	if err != nil {
		panic(err)
	}
	err = productRepository.Add(&mentosMint)
	if err != nil {
		panic(err)
	}

	fmt.Println(productRepository.List())

	mentosFruit.Price.Add(decimal.NewFromFloat(0.01))
	err = productRepository.Save(&mentosFruit)
	if err != nil {
		panic(err)
	}

	err = productRepository.Remove(&mentosMint)
	if err != nil {
		panic(err)
	}

	fmt.Println(productRepository.List())
}
