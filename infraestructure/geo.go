package infraestructure

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/Marcxz/academy-go-q32021/conf"
)

// geo - interface that makes the contract to connects to an external api to do the geocoding process
type Geo interface {
	GeocodingAddress(string) (float64, float64, error)
}

// ig - struct to isolate the infraestructure functions with the repository
type ig struct {
	con *conf.Config
}

// NewGeoInfraestructure - constructor for geo infraestructure, returns a struct with the geo interface requirements
func NewGeoInfraestructure(cfg *conf.Config) Geo {
	return &ig{
		cfg,
	}
}

// GeocodingAddress - function that given an address, connect to an external api and convert it to a lat,lng coordinate, returns the coordinate and an error if exist
func (g *ig) GeocodingAddress(a string) (float64, float64, error) {
	v := make(url.Values)
	v.Add("direccion", a)
	url := fmt.Sprintf("%s?%s", g.con.ApiUrl, v.Encode())

	r, err := http.Get(url)

	if err != nil {
		return -1, -1, err
	}
	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)

	if err != nil {
		return -1, -1, err
	}

	byt := []byte(string(b))
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		return -1, -1, err
	}

	if dat["status"].(float64) != http.StatusOK {
		return -1, -1, errors.New(dat["message"].(string))
	}

	d := dat["data"].(map[string]interface{})
	lat := d["lat"].(float64)
	lng := d["lng"].(float64)

	if lat == -1 && lng == -1 {
		return -1, -1, err
	}
	return lat, lng, nil
}
