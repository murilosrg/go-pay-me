package model

import (
	"github.com/murilosrg/go-pay-me/internal/database"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

//Card struct
type Card struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;" json:"id,omitempty"`
	PAN      string    `gorm:"column:pan;size:14;not null;" json:"pan,omitempty"`
	CVV      string    `gorm:"column:cvv;size:4;not null;" json:"cvv,omitempty"`
	Owner    string    `gorm:"-" json:"owner,omitempty"`
	Validate string    `gorm:"-" json:"validate,omitempty"`
	Flag     string    `gorm:"-" json:"flag,omitempty"`
}

// Create func
func (c *Card) Create() (ra int64, err error) {
	if err = database.DB.Create(&c).Error; err != nil {
		ra = -1
	} else {
		ra = 1
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
func (c *Card) Delete() (ra int64, err error) {
	if err = database.DB.Delete(&c).Error; err != nil {
		ra = -1
	} else {
		ra = 1
	}

	return
}
