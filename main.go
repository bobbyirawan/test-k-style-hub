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
	memberService := service.NewMemberService(db, memberRepo)
	memberController := controller.NewMemberController(memberService)

	route := e.Group("/member")
	{
		route.POST("", memberController.Create)
		route.PUT("/:id", memberController.Update)
		route.DELETE("/:id", memberController.Delete)
		route.GET("", memberController.FindAll)
	}

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
