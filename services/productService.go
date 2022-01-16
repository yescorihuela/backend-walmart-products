package services

import (
	"github.com/yescorihuela/walmart-products/app/response"
	"github.com/yescorihuela/walmart-products/domain"
)

type ProductService interface {
	FindAllProducts() ([]response.ProductResponse, error)
	FindOneProduct(criteria string) (*response.ProductResponse, error)
}

type DefaultProductService struct {
	repo domain.ProductRepository
}

func NewProductService(repo domain.ProductRepository) DefaultProductService {
	return DefaultProductService{repo}
}

func (s DefaultProductService) FindAllProducts() ([]response.ProductResponse, error) {
	products, err := s.repo.GetAllProducts()
	if err != nil {
		return nil, err
	}

	response := domain.ProductToDTOCollection(products)
	return response, nil

}
