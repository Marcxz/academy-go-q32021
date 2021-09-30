package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Marcxz/academy-go-q32021/conf"
	"github.com/Marcxz/academy-go-q32021/usecase"
)

type Reader interface {
	ReadCSVAddress(http.ResponseWriter, *http.Request)
}

type Geocoder interface {
	GeocodeAddress(http.ResponseWriter, *http.Request)
}

type Storer interface {
	StoreGeocodeAddress(http.ResponseWriter, *http.Request)
}

// Address - Interface for Address Controller
type Address interface {
	Reader
	Geocoder
	Storer
}

type ac struct {
	con *conf.Config
	au  usecase.Address
}

// NewAddressController - The constructor for a controller used at routes
func NewAddressController(config *conf.Config, auc usecase.Address) Address {
	return &ac{
		config,
		auc,
	}
}

// ReadCSVAddress - Handler to read the all the Addresses from a csv file
func (c *ac) ReadCSVAddress(w http.ResponseWriter, r *http.Request) {
	ad, err := c.au.ReadCSVAddress(c.con.Filename)

	if err != nil {
		handleError(w, err)
	}

	ja, err := json.Marshal(ad)

	if err != nil {
		handleError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ja)
}

// GeocodeAddress - contoller func to get the address from a query param
func (c *ac) GeocodeAddress(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		handleError(w, err)
	}
	a := r.FormValue("address")

	if len(a) == 0 {
		handleError(w, errors.New("the address should be specified as a queryParam"))
	}

	ga, err := c.au.GeocodeAddress(a)

	if err != nil {
		handleError(w, err)
	}

	jga, err := json.Marshal(ga)

	if err != nil {
		handleError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jga)
}

// StoreGeocodeAddress - geocode an address and store in a csv file
func (c *ac) StoreGeocodeAddress(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		handleError(w, err)
	}
	a := r.FormValue("address")

	if len(a) == 0 {
		handleError(w, errors.New("the address should be specified as a queryParam"))
	}

	ga, err := c.au.StoreGeocodeAddress(a)

	if err != nil {
		handleError(w, err)
	}

	jga, err := json.Marshal(ga)

	if err != nil {
		handleError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jga)
}

// HandleError - Refactored func to report the errors in the controllers
func handleError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
