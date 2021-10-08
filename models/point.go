package models

// Point - struct from a spatial point composed by a lat and lng coordinate
type Point struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
