package services

import (
	"github.com/yescorihuela/walmart-products/app/response"
	"github.com/yescorihuela/walmart-products/domain"
)

type ProductService interface {
	FindAllProducts() ([]response.ProductResponse, error)
	FindProductsByCriteria(string) ([]response.ProductResponse, error)
	FindOneProduct(string) (*response.ProductResponse, error)
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

func (s DefaultProductService) FindProductsByCriteria(criteria string) ([]response.ProductResponse, error) {
	products, err := s.repo.GetProductsByCriteria(criteria)
	if err != nil {
		return nil, err
	}
	response := domain.ProductToDTOCollectionFiltered(products, criteria)
	return response, nil
}

func (s DefaultProductService) FindOneProduct(criteria string) (*response.ProductResponse, error) {
	product, err := s.repo.GetProduct(criteria)
	if err != nil {
		return nil, err
	}
	response := product.ToDTO(criteria)
	return &response, nil
}
