package repositories

import (
	"database_example/core/domain/entities"
	"database_example/core/ports"
	"database_example/infrastructure/database/mappers"
	"database_example/infrastructure/database/orm_entities"
	"errors"

	"gorm.io/gorm"
)

type SqLiteProductRepository struct {
	db *gorm.DB
}

// Add implements ports.ProductRepository
func (s *SqLiteProductRepository) Add(product *entities.Product) error {
	if product == nil {
		return errors.New("product is nil")
	}

	ormProduct := mappers.ProductToORMProduct(*product)
	tx := s.db.Create(&ormProduct)
	if tx.Error != nil {
		return tx.Error
	}

	mappers.UpdateProduct(product, ormProduct)
	return nil
}

// List implements ports.ProductRepository
func (s *SqLiteProductRepository) List() []entities.Product {
	var ormProducts []orm_entities.ORMProduct
	var products []entities.Product
	s.db.Find(&ormProducts)

	for _, ormProduct := range ormProducts {
		products = append(products, mappers.ORMProductToProduct(ormProduct))
	}
	return products
}

// Remove implements ports.ProductRepository
func (s *SqLiteProductRepository) Remove(product *entities.Product) error {
	if product == nil || product.ID == nil {
		return errors.New("please provide a product and product.id")
	}

	ormProduct := mappers.ProductToORMProduct(*product)
	tx := s.db.Delete(&ormProduct)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Save implements ports.ProductRepository
func (s *SqLiteProductRepository) Save(product *entities.Product) error {
	if product == nil || product.ID == nil {
		return errors.New("please provide a product and product.id")
	}

	ormProduct := mappers.ProductToORMProduct(*product)
	tx := s.db.Save(&ormProduct)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

var _ ports.ProductRepository = new(SqLiteProductRepository)

func NewSqLiteProductRepository(db *gorm.DB) *SqLiteProductRepository {
	return &SqLiteProductRepository{
		db: db,
	}
}
