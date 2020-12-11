package main

import (
	"flag"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/murilosrg/go-pay-me/config"
	"github.com/murilosrg/go-pay-me/internal/controller"
)

func main() {
	var shouldInit = flag.Bool("init", false, "initialize data")
	flag.Parse()

	if *shouldInit {
		initAll(config.Config())
	}

	e := echo.New()
	SetupAPIRouter(e)

	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(config.Config().Address))
}

func SetupAPIRouter(e *echo.Echo) {
	e.Use(middleware.Logger())

	group := e.Group("/api")
	{
		group.POST("/pay", controller.CreatePayment)
		group.GET("/cards", controller.GetCard)
		group.POST("/cards", controller.CreateCard)
		group.DELETE("/cards/:id", controller.DeleteCard)
	}
}
