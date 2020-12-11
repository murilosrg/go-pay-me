package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/murilosrg/go-pay-me/internal/model"
	"github.com/murilosrg/go-pay-me/internal/model/request"
	"github.com/murilosrg/go-pay-me/internal/utils"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

// CreatePayment func
func CreatePayment(c echo.Context) error {
	req := request.Payment{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := evaluateCard(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response{
			Message: "card.Token: is invalid token",
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{"nsu": uuid.NewV4().String()})
}

func evaluateCard(req request.Payment) error {
	card := model.Card{ID: req.Card.Token}
	var err error

	if card, err = card.Find(); err != nil {
		return err
	}

	req.Card.PAN = card.PAN
	req.Card.CVV = card.CVV

	return nil
}
