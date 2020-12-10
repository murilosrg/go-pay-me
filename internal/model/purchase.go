package model

//Purchase struct
type Purchase struct {
	Amount       float64  `json:"amount"`
	Installments int      `json:"installments"`
	Items        []string `json:"items"`
}
