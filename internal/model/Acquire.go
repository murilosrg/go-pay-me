package model

//Acquire type
type Acquire string

type acquires struct {
	Stone Acquire
	Cielo Acquire
}

//Acquires acquires list
var Acquires = acquires{
	Stone: "stone",
	Cielo: "cielo",
}
