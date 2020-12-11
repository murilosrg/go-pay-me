package model

import (
	"encoding/json"
	"fmt"
)

//Acquire type
type Acquire int

const (
	Stone = iota
	Cielo
)

var AcquireToString = map[Acquire]string{
	Stone: "Stone",
	Cielo: "Cielo",
}

var AcquireFromString = map[string]Acquire{
	"Stone": Stone,
	"Cielo": Cielo,
}

func (a Acquire) String() string {
	if s, ok := AcquireToString[a]; ok {
		return s
	}
	return "unknown"
}

func (a Acquire) MarshalJSON() ([]byte, error) {
	if s, ok := AcquireToString[a]; ok {
		return json.Marshal(s)
	}
	return nil, fmt.Errorf("unknown acquire %d", a)
}

func (a *Acquire) UnmarshalJSON(text []byte) error {
	var s string
	if err := json.Unmarshal(text, &s); err != nil {
		return err
	}
	var acquire Acquire
	var ok bool
	if acquire, ok = AcquireFromString[s]; !ok {
		return fmt.Errorf("unknown acquire %v", a)
	}
	*a = acquire
	return nil
}
