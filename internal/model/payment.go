package model

// Payment struct
type Payment struct {
	Card     Card     `json:"card"`
	Purchase Purchase `json:"purchase"`
	Seller   Seller   `json:"seller"`
	Acquire  Acquire  `json:"acquire"`
}
