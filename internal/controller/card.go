package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/murilosrg/go-pay-me/internal/model"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

// GetCard func
func GetCard(c echo.Context) error {
	var card model.Card
	var cards []model.Card
	var err error

	if cards, err = card.Get(); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, cards)
}

// CreateCard func
func CreateCard(c echo.Context) error {
	request := model.Card{}

	if err := c.Bind(&request); err != nil {
		return err
	}

	if _, err := request.Create(); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, request)
}

// CreateCard func
func DeleteCard(c echo.Context) error {
	request := model.Card{}
	request.ID =  uuid.FromStringOrNil(c.Param("id"))

	if err := c.Bind(&request); err != nil {
		return err
	}

	if _, err := request.Delete(); err != nil {
		return err
	}

	return c.JSON(http.StatusNoContent, nil)
}