package controller

import (
	"test-k-style-hub/dto"
	"test-k-style-hub/service"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (a *ProductController) FindByProductID(ctx echo.Context) error {
	var (
		res = new(dto.ProductFindByIDResponse)
		req = new(dto.ProductFindByIDRequest)
		err error
	)

	if err := ctx.Bind(req); err != nil {
		return err
	}

	if res, err = a.productService.FindByProductID(req); err != nil {
		return err
	}

	return ctx.JSON(200, res)
}
