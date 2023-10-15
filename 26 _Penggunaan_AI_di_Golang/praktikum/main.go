package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"AI/handlers"
	"AI/usecase"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Cannot load .env file. Err: %s", err)
	}
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	laptopUsecase := usecase.NewLaptopUsecase()
	laptopHandler := handler.NewLaptopHandler(laptopUsecase)
	e.POST("/recommendation-laptop", laptopHandler.RecommendLaptop)
	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	e.Start(port)
}