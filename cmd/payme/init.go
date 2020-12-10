package main

import (
	"context"
	"fmt"
	"github.com/murilosrg/go-pay-me/config"
	"github.com/murilosrg/go-pay-me/internal/database"
	"github.com/murilosrg/go-pay-me/internal/model"
	uuid "github.com/satori/go.uuid"
)

var ctx = context.Background()

func initAll(conf *config.Configuration) {
	if (database.DB.HasTable(&model.Card{})) {
		fmt.Println("db has the table card, so drop it.")
		database.DB.DropTable(&model.Card{})
	}

	database.DB.AutoMigrate(&model.Card{})

	c0 := model.Card{
		ID:  uuid.FromStringOrNil("4a0025ff-3ef1-4fd9-9e1f-26f7ae1acebc"),
		PAN: "5157 2044 8182 7184",
		CVV: "866",
	}

	c1 := model.Card{
		ID:       uuid.FromStringOrNil("cf517e8d-7aa4-4f1f-b94a-6b6668851523"),
		PAN:      "6062 8227 7378 0183",
		CVV:      "537",
	}

	c0.Create()
	c1.Create()
}
