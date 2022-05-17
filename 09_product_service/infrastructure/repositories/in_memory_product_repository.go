package repositories

import (
	"product_service/core/domain"
	"product_service/core/value_objects"
)

type InMemoryProductRepository struct {
	products   domain.Products
	nextFreeId int64
}

func NewInMemoryProductRepository() *InMemoryProductRepository {
	return &InMemoryProductRepository{
		products:   domain.Products{},
		nextFreeId: 1,
	}
}

func (impr InMemoryProductRepository) List() (domain.Products, error) {
	return impr.products, nil
}

func (impr InMemoryProductRepository) Get(id int64) (domain.Product, error) {
	for _, product := range impr.products {
		if product.ID == id {
			return product, nil
		}
	}

	return domain.Product{}, domain.ErrCouldNotFindProduct(id)
}

func (impr InMemoryProductRepository) Update(updateProductParams value_objects.UpdateProductParams) (domain.Product, error) {
	for i, product := range impr.products {
		if product.ID == updateProductParams.ID {
			product.Category = updateProductParams.Category
			product.Name = updateProductParams.Name
			product.Price = updateProductParams.Price
			impr.products[i] = product
			return product, nil
		}
	}
	return domain.Product{}, domain.ErrCouldNotFindProduct(updateProductParams.ID)
}

func (impr *InMemoryProductRepository) Add(product domain.Product) (int64, error) {
	product.ID = impr.nextFreeId
	impr.products = append(impr.products, product)
	impr.nextFreeId++
	return product.ID, nil
}

func (impr *InMemoryProductRepository) Delete(id int64) error {
	for i, product := range impr.products {
		if product.ID == id {
			impr.products = append(impr.products[:i], impr.products[i+1:]...)
			return nil
		}
	}
	return domain.ErrCouldNotFindProduct(id)
}
