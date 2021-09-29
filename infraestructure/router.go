package infraestructure

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Getter interface {
	Get(uri string, f func(http.ResponseWriter, *http.Request))
}

type Poster interface {
	Post(uri string, f func(http.ResponseWriter, *http.Request))
}

type Server interface {
	Serve(p string)
}

// Router - Interface to mock the router interface
type Router interface {
	Getter
	Poster
	Server
}

var (
	md = mux.NewRouter()
)

type muxRouter struct{}

//NewMuxRouter - like the constructor of the Router to handle all the request from the user
func NewMuxRouter() Router {
	return &muxRouter{}
}

/*
func (m *muxRouter) ConfigHandlers() {
	// address Handlers
	m.Get("/address", ac.ReadCSVAddress)
	m.Get("/geocodeAddress", ac.GeocodeAddress)
	m.Get("/storeGeocodeAddress", ac.StoreGeocodeAddress)
}*/

// Get - Refactor and handle the get request from the user
func (*muxRouter) Get(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	md.HandleFunc(uri, f).Methods(http.MethodGet)
}

// Post - Refactor and handle the post request from the user
func (*muxRouter) Post(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	md.HandleFunc(uri, f).Methods(http.MethodPost)
}

// Server - Up and run the project
func (*muxRouter) Serve(p string) {
	fmt.Printf("Server is running on port %s", p)
	http.ListenAndServe(p, md)
}
