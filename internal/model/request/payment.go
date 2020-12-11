package request

import "github.com/murilosrg/go-pay-me/internal/model"

// Payment struct
type Payment struct {
	Card     Card          `json:"card"`
	Purchase Purchase      `json:"purchase"`
	Seller   Seller        `json:"seller"`
	Acquire  model.Acquire `json:"acquire" validate:"min=8"`
}
