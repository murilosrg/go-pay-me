package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/murilosrg/go-pay-me/internal/model"
	"github.com/murilosrg/go-pay-me/internal/utils"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

//Card model
type CardController struct{}

//NewPaymentController creates a new instance
func NewCardController() *CardController {
	return &CardController{}
}

// GetCard func
func (cc *CardController) GetCard(c echo.Context) error {
	var card model.Card
	var cards []model.Card
	var err error

	if cards, err = card.Get(); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, cards)
}

// CreateCard func
func (cc *CardController) CreateCard(c echo.Context) error {
	request := model.Card{}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := request.Create(); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, request)
}

// CreateCard func
func (cc *CardController) DeleteCard(c echo.Context) error {
	request := model.Card{}
	request.ID = uuid.FromStringOrNil(c.Param("id"))

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusNotFound, utils.Response{
			Message: "Card not found",
		})
	}

	if err := request.Delete(); err != nil {
		return c.JSON(http.StatusNotFound, utils.Response{
			Message: "Card not found",
		})
	}

	return c.NoContent(http.StatusNoContent)
}
