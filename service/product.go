package service

import (
	"test-k-style-hub/dto"
	"test-k-style-hub/repository"

	"gorm.io/gorm"
)

type adaptorProductService struct {
	productService repository.ProductRepository
	db             *gorm.DB
}

func NewProductService(productService repository.ProductRepository, db *gorm.DB) ProductService {
	return &adaptorProductService{
		productService: productService,
		db:             db,
	}
}

type ProductService interface {
	FindByProductID(req *dto.ProductFindByIDRequest) (*dto.ProductFindByIDResponse, error)
}

func (a *adaptorProductService) FindByProductID(req *dto.ProductFindByIDRequest) (*dto.ProductFindByIDResponse, error) {
	var (
		product = new([]dto.ProductFindByIDResponseData)
		res     = new(dto.ProductFindByIDResponse)
		err     error
	)

	if product, err = a.productService.FindByID(a.db, req.ProductID); err != nil {
		return nil, err
	}

	// res.Data = product
	res.Message = "request was successfully"
	res.Data = *product

	return res, nil

}
