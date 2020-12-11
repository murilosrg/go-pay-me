package request

import (
	uuid "github.com/satori/go.uuid"
)

//Card struct
type Card struct {
	Token      uuid.UUID `json:"token,omitempty"`
	PAN        string    `json:"pan,omitempty"`
	CVV        string    `json:"cvv,omitempty"`
	Owner      string    `json:"owner,omitempty"`
	Expiration string    `json:"expiration,omitempty"`
	Brand      string    `json:"brand,omitempty"`
}
