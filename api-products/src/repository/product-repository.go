package repository

import (
	"fmt"
	"os"
	"products/src/pb/products"

	"google.golang.org/protobuf/proto"
)

type ProductRepository struct{}

const filename string = "products.txt"

func (pr *ProductRepository) SaveData(productList products.ProductList) error {
	data, err := proto.Marshal(&productList)
	if err != nil {
		return fmt.Errorf("failed to marshal products list: %w", err)
	}
	err = os.WriteFile(filename, data, os.FileMode(0o644))
	return nil
}

func (pr *ProductRepository) FindAll() (products.ProductList, error) {
	productList := products.ProductList{}
	data, err := os.ReadFile(filename)
	if err != nil {
		return productList, fmt.Errorf("failed to read products file: %w", err)
	}
	err = proto.Unmarshal(data, &productList)
	return productList, nil
}

func (pr *ProductRepository) Create(product products.Product) (products.Product, error) {
	productList, err := pr.FindAll()
	if err != nil {
		return product, err
	}
	product.Id = int32(len(productList.Products) + 1)
	productList.Products = append(productList.Products, &product)
	err = pr.SaveData(productList)

	return product, err
}
