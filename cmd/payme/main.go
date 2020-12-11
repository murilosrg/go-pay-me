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

	paymentController := controller.NewPaymentController()
	cardController := controller.NewCardController()

	group := e.Group("/api")
	{
		group.POST("/pay", paymentController.CreatePayment)
		group.GET("/cards", cardController.GetCard)
		group.POST("/cards", cardController.CreateCard)
		group.DELETE("/cards/:id", cardController.DeleteCard)
	}
}
