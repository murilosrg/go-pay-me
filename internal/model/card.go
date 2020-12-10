package model

//Card struct
type Card struct {
	GORMBase
	PAN      string `gorm:"column:pan;size:14;not null;" json:"pan,omitempty"`
	CVV      string `gorm:"column:cvv;size:4;not null;" json:"cvv,omitempty"`
	Owner    string `json:"owner"`
	Validate string `json:"validate"`
	Flag     string `json:"flag"`
}
