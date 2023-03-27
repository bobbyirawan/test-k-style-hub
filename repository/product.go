package repository

import (
	"test-k-style-hub/dto"

	"gorm.io/gorm"
)

type adaptorProductRepository struct {
}

func NewProductRepository() ProductRepository {
	return &adaptorProductRepository{}
}

type ProductRepository interface {
	FindByID(db *gorm.DB, ProductID int) (*[]dto.ProductFindByIDResponseData, error)
}

func (a *adaptorProductRepository) FindByID(db *gorm.DB, ProductID int) (*[]dto.ProductFindByIDResponseData, error) {
	product := new([]dto.ProductFindByIDResponseData)
	// product := new(entity.Products)
	Select := `m.username, 
		m.gender, 
		m.skin_type, 
		m.skin_color, 
		rp.review_desc as desc_review,
		COUNT(lr.review_id) as jumlah_like_review`
	Group := `m.username, 
		m.gender, 
		m.skin_type, 
		m.skin_color, 
		desc_review`

	err := db.Table("products p").Select(Select).
		Joins("LEFT JOIN review_products rp on rp.product_id = p.product_id ").
		Joins("LEFT JOIN members m on m.member_id  = rp.member_id").
		Joins("LEFT JOIN like_reviews lr on lr.review_id = rp.review_id").
		Where("p.product_id = ?", ProductID).
		Group(Group).Scan(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}
