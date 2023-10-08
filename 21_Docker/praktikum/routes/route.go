package routes

import (
	"docker/controllers"
	"docker/repositories"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitRoute(e *echo.Echo, db *gorm.DB) {
	UserRepository := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(*&UserRepository)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format : "[${time_rfc3339}] status=${status} method=${method} uri=${uri} latency=${latency_human} \n",
	}))

	e.GET("/users", userController.GetAllUsers)
	e.POST("/users", userController.CreateUser)
}