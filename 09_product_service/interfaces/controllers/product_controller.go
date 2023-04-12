package controllers

import "database_example/core/ports"

type ProductController struct {
	productRepository *ports.ProductRepository
}

func NewProductController(productRepository *ports.ProductRepository) *ProductController {
	return &ProductController{
		productRepository: productRepository,
	}
}
