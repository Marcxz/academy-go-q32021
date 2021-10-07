package infraestructure

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Getter - interface that makes the contract to Get Handlers
type Getter interface {
	Get(uri string, f func(http.ResponseWriter, *http.Request))
}

// Poster - interface that makes the contract to Post Handlers
type Poster interface {
	Post(uri string, f func(http.ResponseWriter, *http.Request))
}

// Server - interface that makes the contract to create a new api server
type Server interface {
	Serve(p string)
}

// Router - Interface that concatenate all the router interfaces
type Router interface {
	Getter
	Poster
	Server
}

// muxRouter - struct to isolate the infraestructure with repository layer
type muxRouter struct {
	r *mux.Router
}

//NewRouterInfraestructure - router constructor func to create a router struct with the Router interface requirements
func NewRouterInfraestructure(r *mux.Router) Router {
	return &muxRouter{
		r,
	}
}

// Get - refactor and handle the get request from the user
func (m *muxRouter) Get(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	m.r.HandleFunc(uri, f).Methods(http.MethodGet)
}

// Post - refactor and handle the post request from the user
func (m *muxRouter) Post(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	m.r.HandleFunc(uri, f).Methods(http.MethodPost)
}

// Server - func to up and run the project
func (m *muxRouter) Serve(p string) {
	fmt.Printf("Server is running on port %s", p)
	http.ListenAndServe(p, m.r)
}
