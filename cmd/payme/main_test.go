package main

import (
	"github.com/gavv/httpexpect/v2"
	"github.com/labstack/echo/v4"
	"github.com/murilosrg/go-pay-me/config"
	"github.com/murilosrg/go-pay-me/internal/model"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"net/http/httptest"
	"testing"
)

var server *httptest.Server
var card = model.Card{
	ID:  uuid.NewV4(),
	PAN: "1234 1324 1324 1324",
	CVV: "123",
}

func init() {
	initAll(config.Config())
	card.Create()

	// create http.Handler
	e := echo.New()
	SetupAPIRouter(e)

	// run server using httptest
	server = httptest.NewServer(e)
}

// TestCard
func TestCard(t *testing.T) {
	e := httpexpect.New(t, server.URL)

	cardId := uuid.NewV4().String()

	e.POST("/api/cards").
		WithJSON(map[string]interface{}{"id": cardId, "pan": "1234 1234 1234 1234", "cvv": "123"}).
		Expect().
		Status(http.StatusCreated).
		JSON().
		Object().
		ContainsKey("pan").
		ValueEqual("pan", "1234 1234 1234 1234")

	e.GET("/api/cards").
		Expect().
		Status(http.StatusOK).
		JSON().
		Array().
		Element(0).
		Object().
		ContainsKey("pan")

	e.DELETE("/api/cards/" + cardId).
		Expect().
		Status(http.StatusNoContent)
}

// TestPay
func TestPay(t *testing.T) {
	e := httpexpect.New(t, server.URL)

	//Test invalid card id
	e.POST("/api/pay").
		WithJSON(testModels.InvalidPayment).
		Expect().
		Status(http.StatusUnprocessableEntity).
		JSON().
		Object().
		ContainsKey("message").
		ValueEqual("message", "card.Token: is invalid token")

	e.POST("/api/pay").
		WithJSON(testModels.StoneValidPayment).
		Expect().
		Status(http.StatusCreated)

	e.POST("/api/pay").
		WithJSON(testModels.StoneInvalidPayment).
		Expect().
		Status(http.StatusUnprocessableEntity).
		JSON().
		Object().
		ContainsKey("message").
		ValueEqual("message", "Payment unauthorized")

	e.POST("/api/pay").
		WithJSON(testModels.CieloValidPayment).
		Expect().
		Status(http.StatusCreated)

	e.POST("/api/pay").
		WithJSON(testModels.CieloInvalidPayment).
		Expect().
		Status(http.StatusUnprocessableEntity).
		JSON().
		Object().
		ContainsKey("message").
		ValueEqual("message", "Payment unauthorized")
}

type TestModels struct {
	InvalidPayment      map[string]interface{}
	StoneValidPayment   map[string]interface{}
	StoneInvalidPayment map[string]interface{}
	CieloValidPayment   map[string]interface{}
	CieloInvalidPayment map[string]interface{}
}

var testModels = TestModels{
	InvalidPayment: map[string]interface{}{
		"id": "61a2e850-996c-4480-af96-0d93380776e3",
	},
	StoneValidPayment: map[string]interface{}{
		"card": map[string]interface{}{
			"token":      card.ID.String(),
			"owner":      "test owner",
			"expiration": "08/28",
			"brand":      "VISA",
		},
		"purchase": map[string]interface{}{
			"amount":       15.8,
			"installments": 1,
			"Items":        []string{"test", "test2"},
		},
		"seller": map[string]interface{}{
			"document":   "123.132.156-87",
			"address":    "Rua 2",
			"postalCode": "99087-845",
		},
		"acquire": "Stone",
	},
	StoneInvalidPayment: map[string]interface{}{
		"card": map[string]interface{}{
			"token":      card.ID.String(),
			"owner":      "test owner",
			"expiration": "08/28",
			"brand":      "VISA",
		},
		"purchase": map[string]interface{}{
			"amount":       5.8,
			"installments": 1,
			"Items":        []string{"test", "test2"},
		},
		"seller": map[string]interface{}{
			"document":   "123.132.156-87",
			"address":    "Rua 2",
			"postalCode": "99087-845",
		},
		"acquire": "Stone",
	},
	CieloValidPayment: map[string]interface{}{
		"card": map[string]interface{}{
			"token":      card.ID.String(),
			"owner":      "test owner",
			"expiration": "08/28",
			"brand":      "VISA",
		},
		"purchase": map[string]interface{}{
			"amount":       15.8,
			"installments": 1,
			"Items":        []string{"test", "test2"},
		},
		"seller": map[string]interface{}{
			"document":   "123.132.156-87",
			"address":    "Rua 2",
			"postalCode": "99087-845",
		},
		"acquire": "Cielo",
	},
	CieloInvalidPayment: map[string]interface{}{
		"card": map[string]interface{}{
			"token":      card.ID.String(),
			"owner":      "test owner",
			"expiration": "08/28",
			"brand":      "VISA",
		},
		"purchase": map[string]interface{}{
			"amount":       1.8,
			"installments": 1,
			"Items":        []string{"test", "test2"},
		},
		"seller": map[string]interface{}{
			"document":   "123.132.156-87",
			"address":    "Rua 2",
			"postalCode": "99087-845",
		},
		"acquire": "Cielo",
	},
}
