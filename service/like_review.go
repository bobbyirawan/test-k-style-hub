package service

import (
	"test-k-style-hub/dto"
	"test-k-style-hub/entity"
	"test-k-style-hub/repository"

	"gorm.io/gorm"
)

type adaptorLikeReviewService struct {
	likeReviewRepo repository.LikeReviewRepository
	db             *gorm.DB
}

func NewLikeReviewService(likeReviewRepo repository.LikeReviewRepository, db *gorm.DB) LikeReviewService {
	return &adaptorLikeReviewService{
		likeReviewRepo: likeReviewRepo,
		db:             db,
	}
}

type LikeReviewService interface {
	Create(req *dto.LikeReviewCreateRequest) (*dto.LikeReviewCreateResponse, error)
	Delete(req *dto.LikeReviewDeleteRequest) (*dto.LikeReviewDeleteResponse, error)
}

func (a *adaptorLikeReviewService) Create(req *dto.LikeReviewCreateRequest) (*dto.LikeReviewCreateResponse, error) {
	var (
		res        = new(dto.LikeReviewCreateResponse)
		likeReview = new(entity.LikeReviews)
	)

	likeReview.MemberID = req.MemberID
	likeReview.ReviewID = req.ReviewID

	if err := a.likeReviewRepo.Create(a.db, likeReview); err != nil {
		return nil, err
	}

	res.Message = "Like Review Created"

	return res, nil
}

func (a *adaptorLikeReviewService) Delete(req *dto.LikeReviewDeleteRequest) (*dto.LikeReviewDeleteResponse, error) {
	var (
		res        = new(dto.LikeReviewDeleteResponse)
		likeReview = new(entity.LikeReviews)
	)

	likeReview.LikeReviewID = req.LikeReviewID

	if err := a.likeReviewRepo.HardDelete(a.db, likeReview); err != nil {
		return nil, err
	}

	res.Message = "Unlike succeed"

	return res, nil
}
