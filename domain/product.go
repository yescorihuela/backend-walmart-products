package domain

import (
	"strings"

	"github.com/yescorihuela/walmart-products/app/response"
)

type Product struct {
	Id          uint    `json:"id"`
	Brand       string  `json:"brand"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

type ProductRepository interface {
	GetAllProducts() ([]Product, error)
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

func (product Product) ProductDTO(palindromeCriteria string) response.ProductResponse {
	if isPalindrome(palindromeCriteria) {
		return response.ProductResponse{
			Id:          product.Id,
			Brand:       product.Brand,
			Description: product.Description,
			Price:       discount(product.Price, 50),
			HasDiscount: true,
			Discount:    50,
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

func discount(price float32, discount int) float32 {
	return (price * float32(discount)) * 100
}

func isPalindrome(value string) bool {
	result := ""
	for _, s := range strings.ToLower(value) {
		result = result + string(s)
	}
	return strings.ToLower(value) == result
}
