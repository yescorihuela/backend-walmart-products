package domain

import (
	"strings"

	"github.com/yescorihuela/walmart-products/app/response"
	"github.com/yescorihuela/walmart-products/errs"
)

const APPLIED_DISCOUNT = 50

type Product struct {
	Id          uint    `json:"id"`
	Brand       string  `json:"brand"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float32 `json:"price"`
}

type ProductRepository interface {
	GetAllProducts() ([]Product, *errs.AppError)
	GetProductsByCriteria(string) ([]Product, *errs.AppError)
}

func NewProduct(id uint, brand, description, image string, price float32) Product {
	newProduct := Product{
		Id:          id,
		Brand:       brand,
		Description: description,
		Image:       image,
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
			Image:       product.Image,
			Price:       product.Price - discount(product.Price, APPLIED_DISCOUNT),
			HasDiscount: true,
			Discount:    APPLIED_DISCOUNT,
		}
	}
	return response.ProductResponse{
		Id:          product.Id,
		Brand:       product.Brand,
		Description: product.Description,
		Image:       product.Image,
		Price:       product.Price,
		HasDiscount: false,
	}
}

func (product Product) ProductDTORaw() response.ProductResponse {
	return response.ProductResponse{
		Id:          product.Id,
		Brand:       product.Brand,
		Description: product.Description,
		Image:       product.Image,
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
		result = string(s) + result
	}
	return strings.ToLower(value) == result
}
