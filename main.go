package main

import (
	"log"
	"os"
	"test-k-style-hub/controller"
	"test-k-style-hub/di"
	"test-k-style-hub/repository"
	"test-k-style-hub/service"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// environment
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// setup db
	db, err := di.ConnectDB()
	if err != nil {
		panic(err.Error())
	}

	e := echo.New()

	memberRepo := repository.NewMemberRepository()
	productRepo := repository.NewProductRepository()
	likeReviewRepo := repository.NewLikeReviewRepository()

	memberService := service.NewMemberService(db, memberRepo)
	productService := service.NewProductService(productRepo, db)
	likeReviewService := service.NewLikeReviewService(likeReviewRepo, db)

	memberController := controller.NewMemberController(memberService)
	productController := controller.NewProductController(productService)
	LikeReviewController := controller.NewLikeReviewController(likeReviewService)

	member := e.Group("/member")
	{
		member.POST("", memberController.Create)
		member.PUT("/:id", memberController.Update)
		member.DELETE("/:id", memberController.Delete)
		member.GET("", memberController.FindAll)
	}

	product := e.Group("/product")
	{
		product.GET("", productController.FindByProductID)
	}

	like_review := e.Group("/like")
	{
		like_review.POST("", LikeReviewController.Create)
		like_review.DELETE("/:id", LikeReviewController.Delete)
	}

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
