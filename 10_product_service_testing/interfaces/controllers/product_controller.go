package controllers

import (
	"database_example/core/domain/entities"
	"database_example/core/ports"
	"database_example/interfaces/controllers/dtos"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productRepository ports.ProductRepository
}

func (p *ProductController) List(c *gin.Context) {
	products := p.productRepository.List()
	c.JSON(http.StatusOK, products)
}

func parseProductId(c *gin.Context) (uint64, error) {
	productIdParam := c.Param("id")
	return strconv.ParseUint(productIdParam, 10, 64)
}

func (p *ProductController) Add(c *gin.Context) {
	var addProductBody dtos.AddProductBody

	if err := c.ShouldBindJSON(&addProductBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := addProductBody.ToProduct()
	if err := p.productRepository.Add(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// logger.Error(err)
		return
	}

	c.JSON(201, product)
}

func (p *ProductController) Remove(c *gin.Context) {
	productId, err := parseProductId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := p.productRepository.Remove(&entities.Product{
		ID: &productId,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}

func (p *ProductController) Update(c *gin.Context) {
	productId, err := parseProductId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updateProductBody dtos.UpdateProductBody

	if err := c.ShouldBindJSON(&updateProductBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := updateProductBody.ToProduct(productId)
	if err := p.productRepository.Save(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, product)
}

func NewProductController(productRepository ports.ProductRepository) *ProductController {
	return &ProductController{
		productRepository: productRepository,
	}
}
