package routes

import (
	"18_Middleware/praktikum/controllers"
	m "18_Middleware/praktikum/middleware"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	Auth := e.Group("")
	m.LogMiddleware(e) 
	

	e.POST("/users", controllers.CreateUserController)
	e.POST("/users/login", controllers.LoginUserController)
	Auth.GET("/users", controllers.GetUsersController)
	Auth.GET("/users/:id", controllers.GetUserController)
	Auth.DELETE("/users/:id", controllers.DeleteUserController)
	Auth.PUT("/users/:id", controllers.UpdateUserController)

	Auth.GET("/books", controllers.GetBooksController)
	Auth.GET("/books/:id", controllers.GetBookController)
	Auth.POST("/books", controllers.CreateBookController)
	Auth.DELETE("/books/:id", controllers.DeleteBookController)
	Auth.PUT("/books/:id", controllers.UpdateBookController)

	return e
}