package main

import (
	"fmt"
	"github.com/Achmadqizwini/SportKai/config"
	"github.com/Achmadqizwini/SportKai/factory"
	"github.com/Achmadqizwini/SportKai/middlewares"
	"github.com/Achmadqizwini/SportKai/utils/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.GetConfig()
	db := database.InitDB(cfg)

	e := echo.New()

	factory.InitFactory(e, db)

	// middleware
	middlewares.LogMiddlewares(e)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.AppConfig.AppPort)))
}
