package service

import (
	"ecommerceDevFullCycle/internal/database"
	"ecommerceDevFullCycle/internal/entity"
)

type ProductService struct {
	ProductDB database.ProductDB
}

func NewProductService(productDB database.ProductDB) *ProductService {
	return &ProductService{ProductDB: productDB}
}

func (ps *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) CreateProduct(name, description, categoryID, imageURL string, price float64) (*entity.Product, error) {
	product := entity.NewProduct(name, description, categoryID, imageURL, price)
	_, err := ps.ProductDB.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *ProductService) GetProduct(ProductID string) (*entity.Product, error) {
	product, err := ps.ProductDB.GetProduct(ProductID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *ProductService) GetProductsByCategory(categoryID string) ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetProductsByCategory(categoryID)
	if err != nil {
		return nil, err
	}
	return products, nil
}
