package controllers

import (
	"net/http"
	"product_service/core/domain_services"
	"product_service/core/ports"
	"product_service/core/value_objects"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productRepository ports.ProductRepositoryPort
}

func NewProductController(productRepository ports.ProductRepositoryPort) *ProductController {
	return &ProductController{
		productRepository: productRepository,
	}
}

func (pc ProductController) Get(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	product, err := pc.productRepository.Get(id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, product)
}

func (pc ProductController) List(c *gin.Context) {
	products, err := pc.productRepository.List()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, products)
}

func (pc ProductController) Add(c *gin.Context) {
	var productBody value_objects.ProductBody
	if err := c.ShouldBindJSON(&productBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := domain_services.ProductBodyToProduct(productBody)

	products, err := pc.productRepository.Add(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (pc ProductController) Update(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var productBody value_objects.ProductBody
	if err := c.ShouldBindJSON(&productBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateProductParams := domain_services.ProductBodyToUpdateProductParams(id, productBody)

	product, err := pc.productRepository.Update(updateProductParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (pc ProductController) Delete(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = pc.productRepository.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusNoContent)
}
