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

//Stone model
type Stone struct{}

//Stone creates a new instance
func NewStoneService() *Stone {
	return &Stone{}
}

//Authorize request authorization payment
func (s *Stone) Authorize(req request.Payment) (uuid.UUID, error) {
	var endpoint string

	if req.Purchase.Amount > 10 {
		endpoint = "/success"
	} else {
		endpoint = "/failed"
	}

	payload, err := json.Marshal(req)

	if err != nil {
		return uuid.UUID{}, err
	}

	response, err := http.Post(config.Config().Acquires.StoneUrl+endpoint, "", bytes.NewBuffer(payload))

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
