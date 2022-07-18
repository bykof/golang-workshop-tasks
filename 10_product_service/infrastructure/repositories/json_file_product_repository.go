package repositories

import (
	"encoding/json"
	"fmt"
	"os"
	"product_service/core/application"
	"product_service/core/domain"
	"product_service/core/value_objects"
	"sync"
)

type JSONFileProductRepository struct {
	products   domain.Products
	filePath   string
	fileMutex  sync.Mutex
	nextFreeId int64
}

func NewJSONFileProductRepository(filePath string) (*JSONFileProductRepository, error) {
	products, err := application.ReadProducts(filePath)
	if err != nil {
		return nil, err
	}

	return &JSONFileProductRepository{
		filePath:   filePath,
		products:   products,
		nextFreeId: 1,
	}, nil
}

func (jsonfpr *JSONFileProductRepository) List() (domain.Products, error) {
	return jsonfpr.products, nil
}

func (jsonfpr *JSONFileProductRepository) Sync() error {
	jsonfpr.fileMutex.Lock()
	defer jsonfpr.fileMutex.Unlock()
	productFile, err := os.OpenFile(jsonfpr.filePath, os.O_WRONLY, 0744)
	if err != nil {
		return fmt.Errorf("cannot open file: %w", err)
	}
	defer productFile.Close()

	decoder := json.NewEncoder(productFile)
	err = decoder.Encode(jsonfpr.products)
	if err != nil {
		return fmt.Errorf("encoder: %w", err)
	}
	return nil
}

func (jsonfpr *JSONFileProductRepository) Get(id int64) (domain.Product, error) {
	for _, product := range jsonfpr.products {
		if product.ID == id {
			return product, nil
		}
	}

	return domain.Product{}, domain.ErrCouldNotFindProduct(id)
}

func (jsonfpr *JSONFileProductRepository) Update(updateProductParams value_objects.UpdateProductParams) (domain.Product, error) {
	for i, product := range jsonfpr.products {
		if product.ID == updateProductParams.ID {
			product.Category = updateProductParams.Category
			product.Name = updateProductParams.Name
			product.Price = updateProductParams.Price
			jsonfpr.products[i] = product
			err := jsonfpr.Sync()
			if err != nil {
				return product, nil
			}
			return product, nil
		}
	}
	return domain.Product{}, domain.ErrCouldNotFindProduct(updateProductParams.ID)
}

func (jsonfpr *JSONFileProductRepository) Add(product domain.Product) (int64, error) {
	product.ID = jsonfpr.nextFreeId
	jsonfpr.products = append(jsonfpr.products, product)
	err := jsonfpr.Sync()
	if err != nil {
		return 0, err
	}

	jsonfpr.nextFreeId++
	return product.ID, nil
}

func (jsonfpr *JSONFileProductRepository) Delete(id int64) error {
	for i, product := range jsonfpr.products {
		if product.ID == id {
			jsonfpr.products = append(jsonfpr.products[:i], jsonfpr.products[i+1:]...)
			return jsonfpr.Sync()
		}
	}
	return domain.ErrCouldNotFindProduct(id)
}
