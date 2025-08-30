package repository

import (
	"fmt"
	"products/src/mapper"
	"products/src/model"
	"products/src/pb/products"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (pr *ProductRepository) Create(product products.Product) (products.Product, error) {
	gormProduct := mapper.ConvertProtoToGORM(&product)

	result := pr.db.Create(gormProduct)
	if result.Error != nil {
		return product, fmt.Errorf("failed to create product: %w", result.Error)
	}
	return *mapper.ConvertGORMToProto(gormProduct), nil
}

func (pr *ProductRepository) FindAll() (products.ProductList, error) {
	var gormProducts []model.Product

	result := pr.db.Find(&gormProducts)
	if result.Error != nil {
		return products.ProductList{}, fmt.Errorf("failed to find products: %w", result.Error)
	}

	return *mapper.ConvertGORMListToProto(gormProducts), nil
}

func (pr *ProductRepository) FindByID(id int32) (*products.Product, error) {
	var gormProduct model.Product

	result := pr.db.First(&gormProduct, id)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find product with id %d: %w", id, result.Error)
	}

	return mapper.ConvertGORMToProto(&gormProduct), nil
}

func (pr *ProductRepository) Update(product products.Product) (*products.Product, error) {
	gormProduct := mapper.ConvertProtoToGORM(&product)

	result := pr.db.Save(gormProduct)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to update product: %w", result.Error)
	}

	return mapper.ConvertGORMToProto(gormProduct), nil
}

func (pr *ProductRepository) Delete(id int32) error {
	result := pr.db.Delete(&model.Product{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete product with id %d: %w", id, result.Error)
	}
	return nil
}
