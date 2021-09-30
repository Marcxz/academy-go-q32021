//go:generate mockgen -source geo_repository.go -destination mock/geo_repository_mock.go -package mock
package repository

import "github.com/Marcxz/academy-go-q32021/infraestructure"

type Geo interface {
	GeocodeAddress(string) (float64, float64, error)
}

type gr struct {
	geo infraestructure.Geo
}

// NewGeoRepository - constructor func for geo repository
func NewGeoRepository(igeo infraestructure.Geo) Geo {
	return &gr{
		igeo,
	}
}

// GeocodeAddress - func to geocode an address
func (g *gr) GeocodeAddress(a string) (float64, float64, error) {
	return g.geo.GeocodingAddress(a)
}
