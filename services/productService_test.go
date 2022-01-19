package services

import (
	"testing"

	"github.com/golang/mock/gomock"
	realDomain "github.com/yescorihuela/walmart-products/domain"
	mocks "github.com/yescorihuela/walmart-products/mocks/domain"
)

var image string = "https://picsum.photos/200/300"

func TestFindAllProductsService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockProductRepo := mocks.NewMockProductRepository(ctrl)

	serviceProduct := NewProductService(mockProductRepo)

	allProducts := []realDomain.Product{
		{Id: 111, Brand: "Hamilton Beach", Description: "Awesome CoffeeMaker", Price: 1000.00, Image: image},
		{Id: 222, Brand: "Hamilton Beach", Description: "Awesome Iron", Price: 1200.00, Image: image},
		{Id: 333, Brand: "Hamilton Beach", Description: "Awesome Hairdryer", Price: 900.00, Image: image},
	}

	mockProductRepo.EXPECT().GetAllProducts().Return(allProducts, nil)
	productsRequested, _ := serviceProduct.FindAllProducts()

	if len(allProducts) != len(productsRequested) {
		t.Error("Test failed when matching all existing beers")
	}
}

func TestFindProductsByCriteriaService(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockProductRepo := mocks.NewMockProductRepository(ctrl)

	serviceProduct := NewProductService(mockProductRepo)

	searchedProductById := []realDomain.Product{
		{Id: 111, Brand: "Hamilton Beach", Description: "Awesome CoffeeMaker", Price: 1000.00, Image: image},
	}

	criteria := "111"

	mockProductRepo.EXPECT().GetProductsByCriteria(criteria).Return(searchedProductById, nil)
	productsRequested, _ := serviceProduct.FindProductsByCriteria(criteria)

	if len(searchedProductById) != len(productsRequested) {
		t.Error("Test failed when matching all existing beers")
	}
}
