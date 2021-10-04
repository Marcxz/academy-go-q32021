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

type Handler interface {
	ConfigHandlers()
}

// Router - Interface to mock the router interface
type Router interface {
	Getter
	Poster
	Server
}

type muxRouter struct {
	r *mux.Router
}

//NewRouterInfraestructure - like the constructor of the Router to handle all the request from the user
func NewRouterInfraestructure(r *mux.Router) Router {
	return &muxRouter{
		r,
	}
}

/*
func (m *muxRouter) ConfigHandlers() {
	// address Handlers
	m.Get("/address", m.auc.ReadCSVAddress)
	m.Get("/geocodeAddress", m.auc.GeocodeAddress)
	m.Get("/storeGeocodeAddress", m.auc.StoreGeocodeAddress)
}
*/
// Get - Refactor and handle the get request from the user
func (m *muxRouter) Get(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	m.r.HandleFunc(uri, f).Methods(http.MethodGet)
}

// Post - Refactor and handle the post request from the user
func (m *muxRouter) Post(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	m.r.HandleFunc(uri, f).Methods(http.MethodPost)
}

// Server - Up and run the project
func (m *muxRouter) Serve(p string) {
	fmt.Printf("Server is running on port %s", p)
	http.ListenAndServe(p, m.r)
}
