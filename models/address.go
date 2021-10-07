package models

//Address - struct for the address model composed by an id, address and a position (lat, lng)
type Address struct {
	ID int    `json:"id"`
	A  string `json:"a"`
	P  Point  `json:"p"`
}
