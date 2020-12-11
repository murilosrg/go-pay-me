package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/murilosrg/go-pay-me/config"
	"github.com/murilosrg/go-pay-me/internal/model/request"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
)

//Cielo model
type Cielo struct{}

//NewCieloService creates a new instance
func NewCieloService() *Cielo {
	return &Cielo{}
}

//Authorize request authorization payment
func (c *Cielo) Authorize(req request.Payment) (uuid.UUID, error) {
	var endpoint string

	if req.Purchase.Amount > 10 {
		endpoint = "success"
	} else {
		endpoint = "failed"
	}

	payload, err := json.Marshal(req)

	if err != nil {
		return uuid.UUID{}, err
	}

	response, err := http.Post(config.Config().Acquires.CieloUrl+endpoint, "", bytes.NewBuffer(payload))

	if err != nil {
		return uuid.UUID{}, err
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		return uuid.NewV4(), nil
	}

	body, _ := ioutil.ReadAll(response.Body)

	return uuid.UUID{}, errors.New(string(body))
}
