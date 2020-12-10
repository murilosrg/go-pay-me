package model

import (
	"github.com/murilosrg/go-pay-me/internal/database"
	uuid "github.com/satori/go.uuid"
)

//Card struct
type Card struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;json:omitempty"`
	PAN      string    `gorm:"column:pan;size:14;not null;" json:"pan,omitempty"`
	CVV      string    `gorm:"column:cvv;size:4;not null;" json:"cvv,omitempty"`
	Owner    string    `gorm:"-" json:"owner"`
	Validate string    `gorm:"-" json:"validate"`
	Flag     string    `gorm:"-" json:"flag"`
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