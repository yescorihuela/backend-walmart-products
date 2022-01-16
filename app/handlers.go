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
		ctx.JSON(http.StatusInternalServerError, nil)
	}
	ctx.JSON(http.StatusOK, response)
}
