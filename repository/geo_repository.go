//go:generate mockgen -source geo_repository.go -destination mock/geo_repository_mock.go -package mock
package repository

import "github.com/Marcxz/academy-go-q32021/infraestructure"

type Geo interface {
	GeocodeAddress(string) (float64, float64, error)
}

var ig infraestructure.Geo

type g struct{}

// NewGeoRepository - constructor func for geo repository
func NewGeoRepository(igeo infraestructure.Geo) Geo {
	ig = igeo
	return &g{}
}

// GeocodeAddress - func to geocode an address
func (*g) GeocodeAddress(a string) (float64, float64, error) {
	return ig.GeocodingAddress(a)
}
