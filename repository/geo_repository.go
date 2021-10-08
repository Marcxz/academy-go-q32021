//go:generate mockgen -source geo_repository.go -destination mock/geo_repository_mock.go -package mock
package repository

import "github.com/Marcxz/academy-go-q32021/infraestructure"

// Geo - interface that makes the contract to geocode an address
type Geo interface {
	GeocodeAddress(string) (float64, float64, error)
}

// gr - struct to isolate geo repository funcs from the other layers
type gr struct {
	geo infraestructure.Geo
}

// NewGeoRepository - constructor func to link the geo repository with the usecase
func NewGeoRepository(igeo infraestructure.Geo) Geo {
	return &gr{
		igeo,
	}
}

// GeocodeAddress - func to geocode an address, retuns lat, lng float64 coordinate and an error if exist
func (g *gr) GeocodeAddress(a string) (float64, float64, error) {
	return g.geo.GeocodingAddress(a)
}
