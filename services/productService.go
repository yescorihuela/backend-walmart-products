package services

import (
	"github.com/yescorihuela/walmart-products/app/response"
	"github.com/yescorihuela/walmart-products/domain"
	"github.com/yescorihuela/walmart-products/errs"
)

type ProductService interface {
	FindAllProducts() ([]response.ProductResponse, *errs.AppError)
	FindProductsByCriteria(string) ([]response.ProductResponse, *errs.AppError)
}

type DefaultProductService struct {
	repo domain.ProductRepository
}

func NewProductService(repo domain.ProductRepository) DefaultProductService {
	return DefaultProductService{repo}
}

func (s DefaultProductService) FindAllProducts() ([]response.ProductResponse, *errs.AppError) {
	products, err := s.repo.GetAllProducts()
	if err != nil {
		return nil, err
	}

	response := domain.ProductToDTOCollection(products)
	return response, nil
}

func (s DefaultProductService) FindProductsByCriteria(criteria string) ([]response.ProductResponse, *errs.AppError) {
	products, err := s.repo.GetProductsByCriteria(criteria)
	if err != nil {
		return nil, err
	}
	response := domain.ProductToDTOCollectionFiltered(products, criteria)
	return response, nil
}
