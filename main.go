package main

import (
	"os"

	"github.com/Marcxz/academy-go-q32021/conf"
	"github.com/Marcxz/academy-go-q32021/controller"
	"github.com/Marcxz/academy-go-q32021/infraestructure"
)

var (
	ac = controller.NewAddressController()
	hr = infraestructure.NewMuxRouter()
)

func main() {
	conf.ConfigInit()
	// Handlers for address
	hr.Get("/address", ac.ReadCSVAddress)
	hr.Get("/geocodeAddress", ac.GeocodeAddress)
	hr.Get("/storeGeocodeAddress", ac.StoreGeocodeAddress)
	hr.Serve(os.Getenv("p"))
}
