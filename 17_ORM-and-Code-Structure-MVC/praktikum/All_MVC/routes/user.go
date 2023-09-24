package routes

import (
	"17_ORM-and-Code-Structure-MVC/praktikum/All_MVC/controllers"

	"github.com/labstack/echo"
)

func SetUserRoutes(e *echo.Echo) {
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)
}