package model

import (
	"github.com/murilosrg/go-pay-me/internal/database"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

//Card struct
type Card struct {
	ID  uuid.UUID `gorm:"type:uuid;primary_key;" json:"id,omitempty"`
	PAN string    `gorm:"column:pan;size:14;not null;" json:"pan,omitempty"`
	CVV string    `gorm:"column:cvv;size:4;not null;" json:"cvv,omitempty"`
}

// Create func
func (c *Card) Create() (err error) {
	if err = database.DB.Create(&c).Error; err != nil {
		logrus.Error(err)
	}

	return
}

// Get func
func (c *Card) Get() (cards []Card, err error) {
	if err = database.DB.Where(&c).Find(&cards).Error; err != nil {
		logrus.Error(err)
	}

	return
}

// Delete func
func (c *Card) Delete() (err error) {
	if err = database.DB.Delete(&c).Error; err != nil {
		logrus.Error(err)
	}

	return
}

// Find func
func (c *Card) Find() (card Card, err error) {
	if err = database.DB.Find(&card, "id = ?", c.ID).Error; err != nil {
		logrus.Error(err)
	}

	return
}
