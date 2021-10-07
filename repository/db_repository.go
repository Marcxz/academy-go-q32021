//go:generate mockgen -source db_repository.go -destination mock/db_repository_mock.go -package mock

package repository

import "github.com/Marcxz/academy-go-q32021/infraestructure"

// RouterGenerator - interface that makes the contract to send to the infraestructure the 2 coordinates to generate a route
type RouteGenerator interface {
	GenerateRoute(float64, float64, float64, float64) (string, error)
}

// GeoDB - interface which concatenates all the db repository interfaces
type GeoDB interface {
	RouteGenerator
}

// geoDBR - struct to isolate the repository db layer with the rest
type geoDBR struct {
	igeodb infraestructure.GeoDB
}

// NewGeoDBRepository - constructor func to link the db repository with usecase layer, the db repository accomplish the GeoDB requirements
func NewGeoDBRepository(igeodb infraestructure.GeoDB) GeoDB {
	return &geoDBR{
		igeodb,
	}
}

// GenerateRoute - func to send to db infraestructure 2 coordinates to generate a route, retuns a geojson string route and an error if exist
func (geor *geoDBR) GenerateRoute(latA float64, lngA float64, latB float64, lngB float64) (string, error) {
	rp, err := geor.igeodb.GenerateRoute(latA, lngA, latB, lngB)

	if err != nil {
		return rp, err
	}

	return rp, nil
}
