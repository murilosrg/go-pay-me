package controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/murilosrg/go-pay-me/internal/model"
	"github.com/murilosrg/go-pay-me/internal/model/request"
	"github.com/murilosrg/go-pay-me/internal/service"
	"github.com/murilosrg/go-pay-me/internal/utils"
	"net/http"
)

//PaymentController model
type PaymentController struct {
	Handler *service.AuthorizationHandler
}

//NewPaymentController creates a new instance
func NewPaymentController() *PaymentController {
	return &PaymentController{
		Handler: service.NewAuthorizationHandler(),
	}
}

// CreatePayment func
func (p *PaymentController) CreatePayment(c echo.Context) error {
	req := request.Payment{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := evaluateCard(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response{
			Message: "card.Token: is invalid token",
		})
	}

	nsu, err := p.Handler.Handle(req)

	if err != nil {
		resp := utils.Response{}
		json.Unmarshal([]byte(err.Error()), &resp)

		return c.JSON(http.StatusUnprocessableEntity, resp)
	}

	return c.JSON(http.StatusCreated, map[string]string{"nsu": nsu.String()})
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
