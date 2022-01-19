package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yescorihuela/walmart-products/services"
)

type ProductHandlers struct {
	productService services.ProductService
}

func (ph *ProductHandlers) GetAllProducts(ctx *gin.Context) {
	response, err := ph.productService.FindAllProducts()
	if err != nil {
		ctx.JSON(err.Code, nil)
	}
	ctx.JSON(http.StatusOK, response)
}

func (ph *ProductHandlers) SearchByCriteria(ctx *gin.Context) {
	criteria := ctx.Query("q")
	response, err := ph.productService.FindProductsByCriteria(criteria)
	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
	}
	ctx.JSON(http.StatusOK, response)
}
