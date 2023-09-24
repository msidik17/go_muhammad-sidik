package routes

import (
	"17_ORM-and-Code-Structure-MVC/praktikum/All_MVC/controllers"

	"github.com/labstack/echo"
)

func SetBlogRoutes(e *echo.Echo) {
	e.GET("/blogs", controllers.GetBlogsController)
	e.GET("/blogs/:id", controllers.GetBlogController)
	e.POST("/blogs", controllers.CreateBlogController)
	e.DELETE("/blogs/:id", controllers.DeleteBlogController)
	e.PUT("/blogs/:id", controllers.UpdateBlogController)
}