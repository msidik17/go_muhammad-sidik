package main

import (
	"17_ORM-and-Code-Structure-MVC/praktikum/All_MVC/configs"
	"17_ORM-and-Code-Structure-MVC/praktikum/All_MVC/routes"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

func main() {
	configs.InitDB()
	e := echo.New()
	routes.SetUserRoutes(e)
	routes.SetBookRoutes(e)
	routes.SetBlogRoutes(e)
	e.Logger.Fatal(e.Start(":8000"))
}

