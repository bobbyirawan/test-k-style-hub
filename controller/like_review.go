package controller

import (
	"test-k-style-hub/dto"
	"test-k-style-hub/service"

	"github.com/labstack/echo/v4"
)

type LikeReviewController struct {
	likeReviewService service.LikeReviewService
}

func NewLikeReviewController(likeReviewService service.LikeReviewService) *LikeReviewController {
	return &LikeReviewController{
		likeReviewService: likeReviewService,
	}
}

func (l *LikeReviewController) Create(ctx echo.Context) error {
	var (
		req = new(dto.LikeReviewCreateRequest)
	)

	if err := ctx.Bind(req); err != nil {
		return err
	}

	res, err := l.likeReviewService.Create(req)
	if err != nil {
		return err
	}

	return ctx.JSON(201, res)
}

func (l *LikeReviewController) Delete(ctx echo.Context) error {
	var (
		req = new(dto.LikeReviewDeleteRequest)
	)

	if err := ctx.Bind(req); err != nil {
		return err
	}

	res, err := l.likeReviewService.Delete(req)
	if err != nil {
		return err
	}

	return ctx.JSON(200, res)
}
