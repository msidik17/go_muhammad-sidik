package main

import (
	"18_Middleware/praktikum/configs"
	"18_Middleware/praktikum/routes"
)

func main() {
	configs.InitDB()
	
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}


