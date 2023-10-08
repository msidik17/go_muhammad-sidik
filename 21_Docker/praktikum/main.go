package main

import (
	"docker/configs"
	"docker/databases"
	"docker/routes"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	conf := configs.InitConfig()
	db := databases.InitDBMysql(conf)
	configs.InitMigration(db)

	e := echo.New()
	
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	routes.InitRoute(e, db)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", conf.SERVER_PORT)))
}