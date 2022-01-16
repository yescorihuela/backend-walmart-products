package domain

import (
	"strings"

	"github.com/yescorihuela/walmart-products/app/response"
)

const APPLIED_DISCOUNT = 50

type Product struct {
	Id          uint    `json:"id"`
	Brand       string  `json:"brand"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

type ProductRepository interface {
	GetAllProducts() ([]Product, error)
	GetProductsByCriteria(string) ([]Product, error)
	GetProduct(string) (*Product, error)
}

func NewProduct(id uint, brand, description string, price float32) Product {
	newProduct := Product{
		Id:          id,
		Brand:       brand,
		Description: description,
		Price:       price,
	}
	return newProduct
}

func (product Product) ToDTO(criteria string) response.ProductResponse {
	if isPalindrome(criteria) && criteria != "" {
		return response.ProductResponse{
			Id:          product.Id,
			Brand:       product.Brand,
			Description: product.Description,
			Price:       product.Price - discount(product.Price, APPLIED_DISCOUNT),
			HasDiscount: true,
			Discount:    APPLIED_DISCOUNT,
		}
	}
	return response.ProductResponse{
		Id:          product.Id,
		Brand:       product.Brand,
		Description: product.Description,
		Price:       product.Price,
		HasDiscount: false,
	}
}

func (product Product) ProductDTORaw() response.ProductResponse {
	return response.ProductResponse{
		Id:          product.Id,
		Brand:       product.Brand,
		Description: product.Description,
		Price:       product.Price,
		HasDiscount: false,
	}
}

func ProductToDTOCollection(products []Product) []response.ProductResponse {
	var productsCollection []response.ProductResponse
	for _, product := range products {
		productsCollection = append(productsCollection, product.ProductDTORaw())
	}
	return productsCollection
}

func ProductToDTOCollectionFiltered(products []Product, criteria string) []response.ProductResponse {
	var productsCollection []response.ProductResponse
	for _, product := range products {
		productsCollection = append(productsCollection, product.ToDTO(criteria))
	}
	return productsCollection
}

func discount(price float32, discount uint) float32 {
	return (price * float32(discount)) / 100
}

func isPalindrome(value string) bool {
	result := ""
	for _, s := range strings.ToLower(value) {
		result = result + string(s)
	}
	return strings.ToLower(value) == result
}
