package app

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/yescorihuela/walmart-products/app/response"
	"github.com/yescorihuela/walmart-products/errs"
	mocks "github.com/yescorihuela/walmart-products/mocks/services"
)

var image string = "https://picsum.photos/200/300"

func TestGetAllProductsHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mocks.NewMockProductService(ctrl)
	dummyProducts := []response.ProductResponse{
		{Id: 1, Brand: "Hamilton Beach", Description: "The coffeemaker ever", Image: image, Price: 2000.0, HasDiscount: false, Discount: 0},
		{Id: 2, Brand: "Oster", Description: "Another coffeemaker", Image: image, Price: 2210.0, HasDiscount: false, Discount: 0},
		{Id: 3, Brand: "Molinex", Description: "Mid-range coffeemaker", Image: image, Price: 1600.0, HasDiscount: false, Discount: 0},
		{Id: 4, Brand: "Starbucks", Description: "Regular coffeemaker", Image: image, Price: 1900.0, HasDiscount: false, Discount: 0},
	}
	mockService.EXPECT().FindAllProducts().Return(dummyProducts, nil)
	ph := ProductHandlers{productService: mockService}

	r := gin.Default()
	r.GET("/products", ph.GetAllProducts)

	request, _ := http.NewRequest(http.MethodGet, "/products", nil)
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}

}

func TestGetAllProductsNotFoundHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mocks.NewMockProductService(ctrl)

	dummyError := errs.NewNotFoundError("Elements not found")

	mockService.EXPECT().FindAllProducts().Return(nil, dummyError)
	ph := ProductHandlers{productService: mockService}
	r := gin.Default()
	r.GET("/products", ph.GetAllProducts)

	request, _ := http.NewRequest(http.MethodGet, "/products", nil)
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusNotFound {
		t.Error("Failed while testing the status code")
	}
}

func TestSearchByCriteriaHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mocks.NewMockProductService(ctrl)
	dummyProducts := []response.ProductResponse{
		{Id: 111, Brand: "Hamilton Beach", Description: "The coffeemaker ever", Image: image, Price: 1000.0, HasDiscount: false, Discount: 0},
	}
	mockService.EXPECT().FindAllProducts().Return(dummyProducts, nil)
	ph := ProductHandlers{productService: mockService}

	r := gin.Default()
	r.GET("/products/search", ph.GetAllProducts)
	url := fmt.Sprintf("/products/search?q=%v", 111)
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}
