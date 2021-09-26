package repository

import "github.com/Marcxz/academy-go-q32021/infraestructure"

type geo interface {
	GeoCodeAddress(string) (float64, float64, error)
}

type g struct{}

var (
	gi = infraestructure.NewGeoInfraestructure()
)

func NewGeoRepository() geo {
	return &g{}
}

func (*g) GeoCodeAddress(a string) (float64, float64, error) {
	return gi.GeocodingAddress(a)
}
