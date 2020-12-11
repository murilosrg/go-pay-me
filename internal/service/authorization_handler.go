package service

import (
	"github.com/murilosrg/go-pay-me/internal/model"
	"github.com/murilosrg/go-pay-me/internal/model/request"
	uuid "github.com/satori/go.uuid"
)

//AuthorizationHandler model
type AuthorizationHandler struct{
	stoneHandler *Stone
	cieloHandler *Cielo
}

//NewAuthorizationHandler creates a new instance
func NewAuthorizationHandler() *AuthorizationHandler {
	return &AuthorizationHandler{
		stoneHandler: NewStoneService(),
		cieloHandler: NewCieloService(),
	}
}

func (a *AuthorizationHandler) Handle(req request.Payment) (uuid.UUID, error) {
	switch req.Acquire {
	case model.Stone:
		return a.stoneHandler.Authorize(req)
	case model.Cielo:
		return a.cieloHandler.Authorize(req)
	default:
		return uuid.UUID{}, nil
	}
}