package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"text/template"

	"github.com/Marcxz/academy-go-q32021/conf"
	"github.com/Marcxz/academy-go-q32021/usecase"
)

// Reader - interface that makes the contract to handle the read process into a csv address file
type Reader interface {
	ReadCSVAddress(http.ResponseWriter, *http.Request)
}

// Geocoder - interface that makes the contract to handle the geocode process
type Geocoder interface {
	GeocodeAddress(http.ResponseWriter, *http.Request)
}

// Storer - interface that makes the contract to handle the store process to push an address into a csv file
type Storer interface {
	StoreGeocodeAddress(http.ResponseWriter, *http.Request)
}

// Concurrencier - interface that makes the contract to handle the concurrency process to read a csv address with concurrency
type Concurrencier interface {
	ReadAddressCSVConcurrency(http.ResponseWriter, *http.Request)
}

// Router - interface that makes the contract to handle the route process to generate the route from 2 address
type Router interface {
	GenerateRouterFrom2Address(http.ResponseWriter, *http.Request)
}

// Address - interface for address controller
type Address interface {
	Reader
	Geocoder
	Storer
	Concurrencier
	Router
}

// ac - struct to isolate the controller funcs layer
type ac struct {
	con *conf.Config
	au  usecase.Address
}

// NewAddressController - the constructor for a controller accomplishing the Address interface requirements
func NewAddressController(config *conf.Config, auc usecase.Address) Address {
	return &ac{
		config,
		auc,
	}
}

// ReadCSVAddress - handler to read all addresses from a csv file
func (c *ac) ReadCSVAddress(w http.ResponseWriter, r *http.Request) {
	ad, err := c.au.ReadCSVAddress()

	if err != nil {
		handleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ja, err := json.Marshal(struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
	}{http.StatusOK, ad})

	if err != nil {
		handleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ja)
}

// GeocodeAddress - handler to get the address from a query param and init the geocode process
func (c *ac) GeocodeAddress(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		handleError(w, http.StatusBadRequest, err.Error())
		return
	}
	a := r.FormValue("address")

	if len(a) == 0 {
		handleError(w, http.StatusBadRequest, "the address should be sended has a queryParam")
		return
	}

	ga, err := c.au.GeocodeAddress(a)

	if err != nil {
		handleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	jga, err := json.Marshal(struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
	}{http.StatusOK, ga})

	if err != nil {
		handleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jga)
}

// StoreGeocodeAddress - handler to geocode an address and store it in a csv file
func (c *ac) StoreGeocodeAddress(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		handleError(w, http.StatusInternalServerError, err.Error())
		return
	}
	a := r.FormValue("address")

	if len(a) == 0 {
		handleError(w, http.StatusBadRequest, "the address should be specified as a queryParam")
		return
	}

	ga, err := c.au.StoreGeocodeAddress(a)

	if err != nil {
		handleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	jga, err := json.Marshal(struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
	}{http.StatusOK, ga})

	if err != nil {
		handleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jga)
}

// ReadAddressCSVConcurrency - handler to read a csv address file with concurrency, the handler get type (even, odd), #items and #items per worker
func (c *ac) ReadAddressCSVConcurrency(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		handleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	t := r.FormValue("type")
	if len(t) == 0 {
		handleError(w, http.StatusBadRequest, "the type should be passed has a queryparam")
		return
	}

	i := r.FormValue("items")
	if len(i) == 0 {
		handleError(w, http.StatusBadRequest, "should specify the number of items that you want has a queryparam")
		return
	}

	ii, err := strconv.Atoi(i)

	if err != nil {
		handleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ipw := r.FormValue("items_per_worker")
	if len(ipw) == 0 {
		handleError(w, http.StatusBadRequest, "should specify the number of items per worker has a queryparam")
		return
	}

	iipw, err := strconv.Atoi(ipw)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	m, err := c.au.ReadAddressCSVConcurrency(t, ii, iipw)

	if err != nil {
		handleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	jm, err := json.Marshal(struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
	}{http.StatusOK, m})

	if err != nil {
		handleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jm)
}

// GenerateRouterFrom2Addresss - handler to generate the route from 2 Address, takes 2 address, geocode each one, generate the route and render it into a map
func (c *ac) GenerateRouterFrom2Address(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		handleError(w, http.StatusInternalServerError, err.Error())
		return
	}
	f := r.FormValue("from")
	if len(f) <= 0 {
		handleError(w, http.StatusBadRequest, "should specify the from address that you want has a queryparam")
		return
	}

	t := r.FormValue("to")
	if len(t) <= 0 {
		handleError(w, http.StatusBadRequest, "should specify the from address that you want has a queryparam")
		return
	}
	route, err := c.au.GenerateRouterFrom2Address(f, t)

	if err != nil {
		handleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	tpl := template.Must(template.ParseGlob(c.con.MapPath))

	err = tpl.ExecuteTemplate(w, c.con.MapFile, route)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

// HandleError - refactored func to report the errors in the controller funcs
func handleError(w http.ResponseWriter, status int, err string) {
	bytes, _ := json.Marshal(struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}{status, err})

	w.Header().Add("Content-Type", "application/json")
	w.Write(bytes)
}
