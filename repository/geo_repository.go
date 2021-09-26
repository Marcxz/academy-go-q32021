package repository

import "github.com/Marcxz/academy-go-q32021/infraestructure"

type geo interface {
	GeocodeAddress(string) (float64, float64, error)
}

type g struct{}

var gi = infraestructure.NewGeoInfraestructure()

// NewGeoRepository - constructor func for geo repository
func NewGeoRepository() geo {
	return &g{}
}

// GeocodeAddress - func to geocode an address
func (*g) GeocodeAddress(a string) (float64, float64, error) {
	return gi.GeocodingAddress(a)
}
