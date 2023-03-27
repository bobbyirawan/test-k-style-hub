package repository

import (
	"test-k-style-hub/entity"

	"gorm.io/gorm"
)

type adaptorLikeReviewRepository struct {
}

func NewLikeReviewRepository() LikeReviewRepository {
	return &adaptorLikeReviewRepository{}
}

type LikeReviewRepository interface {
	Create(db *gorm.DB, data *entity.LikeReviews) error
	HardDelete(db *gorm.DB, data *entity.LikeReviews) error
}

func (a *adaptorLikeReviewRepository) Create(db *gorm.DB, data *entity.LikeReviews) error {
	if err := db.Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (a *adaptorLikeReviewRepository) HardDelete(db *gorm.DB, data *entity.LikeReviews) error {
	if err := db.Unscoped().Delete(data).Error; err != nil {
		return err
	}
	return nil
}
