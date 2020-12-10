package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/murilosrg/go-pay-me/internal/model"
	"net/http"
)

// CreatePayment func
func CreatePayment(c echo.Context) error {
	request := model.Payment{}

	if err := c.Bind(&request); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, request)
}